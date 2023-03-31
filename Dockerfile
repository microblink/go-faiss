# syntax=docker/dockerfile:1.3
FROM golang:1.20-bullseye
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends libopenblas-dev cmake
RUN git clone https://github.com/facebookresearch/faiss.git /faiss \
  && cd /faiss \
  && git checkout v1.7.3 \
  && cmake -B build -DFAISS_ENABLE_GPU=OFF -DFAISS_ENABLE_C_API=ON -DBUILD_SHARED_LIBS=ON -DFAISS_ENABLE_PYTHON=OFF -DBUILD_TESTING=OFF . \
  && make -C build \
  && make -C build install \
  && cp /faiss/build/c_api/libfaiss_c.so /usr/lib/libfaiss_c.so \
  && rm -rf /faiss
ENV LD_LIBRARY_PATH $LD_LIBRARY_PATH:/usr/lib:/usr/local/lib/
WORKDIR /src
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build