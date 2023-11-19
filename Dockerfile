FROM debian:stable
RUN apt update
RUN apt-get update
RUN apt-get install -y \
    vim \
    golang \
    gcc \
    git \
    ca-certificates \
    libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev  libgl1-mesa-dev xorg-dev \
    libgl1-mesa-dev \
    libxxf86vm-dev \
    dbus-x11 \
    openssh-server \
    xauth
CMD ["bash"]
# winpty docker run --rm -it --net=host -e DISPLAY=host.docker.internal:0 --mount src="$(pwd)",target=/home/dev/,type=bind lfyne
