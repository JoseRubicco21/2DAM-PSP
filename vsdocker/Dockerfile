FROM debian:bookworm

RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    apt-transport-https \
    ca-certificates \
    gnupg2 \
    software-properties-common \
    libasound2 \
    libgtk-3-0 \
    libxss1 \
    libnss3 \
    xvfb \
    libgconf-2-4 \
    libdrm2 \
    libxkbfile1 \
    libxrandr2 \
    libasound2-dev \
    libatspi2.0-0 \
    libdrm2 \
    libgtk-3-0 \
    libxss1

# Code installation
COPY code.deb /tmp/code.deb
RUN dpkg -i /tmp/code.deb || apt-get -f install -y
RUN rm /tmp/code.deb

# GO
COPY go1.25.1.linux-amd64.tar.gz /tmp/go1.25.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go1.25.1.linux-amd64.tar.gz
RUN rm /tmp/go1.25.1.linux-amd64.tar.gz

# Set PATH environment variable permanently
ENV PATH=$PATH:/usr/local/go/bin

# Create the directory for the workspace
RUN mkdir -p /code/
WORKDIR /code/
RUN mkdir -p /.code_user/

# Start with CMD
CMD ["tail", "-f", "/dev/null"]
#ENTRYPOINT ["code", ".", "--no-sandbox", "--user-data-dir", "/code/"]

