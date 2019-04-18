FROM alpine:3.9

RUN apk update && \
    apk add --no-cache ffmpeg=4.0.2-r0 openssl && \
    ffmpeg -h && \
    mkdir -p /var/vod && \
    mkdir /var/input

WORKDIR /var/vod

ENV BASE_URL=https://storage.googleapis.com/vod-m3u8/

COPY ./create_key_info_file.sh /bin/create_key_info_file.sh
RUN chmod +x /bin/create_key_info_file.sh

CMD /bin/create_key_info_file.sh && ffmpeg -i /var/input/video -hls_time 10 -hls_key_info_file file.keyinfo -hls_playlist_type vod -hls_base_url "$BASE_URL$FILE_NAME/" -hls_segment_filename "seg-%d.ts" hls.m3u8