FROM public.ecr.aws/lambda/provided:al2 as build
ENV LD_LIBRARY_PATH="/opt/:$LD_LIBRARY_PATH"
# install compiler
RUN yum install -y golang
RUN yum -y install wget unzip libX11 nano wget unzip xorg-x11-xauth xclock xterm
RUN go env -w GOPROXY=direct
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# build
ADD . .
RUN go build -o /main
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
RUN yum -y install wget unzip libX11 nano wget unzip xorg-x11-xauth xclock xterm
RUN mkdir bin
COPY --from=build /main /main
COPY headless-chromium /opt/headless-chromium
RUN chmod 777 /opt/headless-chromium
ENTRYPOINT [ "/main" ]
