wc="netkit-wc"
image_name="netkit-deb-test"

genrule(
    name="install_packages",
    srcs=["./packages-list", "debconf-package-selections", "packages-custom", ":create_buildah_wc"],
    cmd=[
        # f"buildah config --env DEBIAN_FRONTEND=\"noninteractive\" {wc}",
        f"buildah run --env DEBIAN_FRONTEND=noninteractive {wc} apt update --assume-yes",
        f"cat debconf-package-selections | buildah run {wc} debconf-set-selections",
        f"buildah run --env DEBIAN_FRONTEND=noninteractive {wc} apt install --assume-yes $(cat packages-list | grep -v '#')",
        f"cat packages-custom | buildah run {wc} bash -"
    ]
)

genrule(
    name="copy_fs_tweaks",
    srcs=["filesystem-tweaks", "HOME"],
    deps=[":install_packages"],
    cmd=[
        f"buildah run {wc} useradd netkit -m -s /bin/bash -u 1000 -p $(openssl password -crypt netkit) -G sudo",
        f"buildah copy {wc} filesystem-tweaks /",
        f"buildah run {wc} mkdir -p /root /home/netkit",
        f"buildah copy {wc} HOME /root",
        f"buildah copy {wc} HOME /home/netkit",
    ]
)

genrule(
    name="setup_systemd",
    srcs=["disabled-services"],
    deps=[":copy_fs_tweaks"],
    cmd=[
        f"buildah run {wc} systemctl enable netkit-startup-phase1.service",
        f"buildah run {wc} systemctl enable netkit-startup-phase2.service",
        f"buildah run {wc} systemctl enable netkit-shutdown.service",
        f"buildah run {wc} ln -s /lib/systemd/system/getty@.service /etc/systemd/system/getty.target.wants/getty@tty0.service",
        f"buildah run {wc} systemctl mask getty-static",
        "for i in {2..6}; do buildah run "+ wc + " systemctl mask getty@tty${i}.service; done",
        "for SERVICE in $(cat disabled-services); do buildah run " + wc + " systemctl disable ${SERVICE}; done",
    ]
)

genrule(
    name="create_buildah_wc",
    cmd=[
        f"buildah from --name {wc} --dns none docker.io/library/debian:10",
        f"buildah config --cmd /sbin/init {wc}",
    ],
    outs=["containers-user-1000"],
)

genrule(
    name="oci_image",
    deps=[":create_buildah_wc", ":install_packages", ":copy_fs_tweaks", ":setup_systemd"],
    tools=["buildah"],
    cmd=[
        f"buildah commit {wc} {image_name}",
        f"buildah push {image_name} oci-archive:netkit.oci.tar"
    ],
    outs=["netkit.oci.tar"]
)
