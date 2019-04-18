# Segment Process

Build the image.
`docker build -t hls-ffmpeg .`

Usage:
`docker run -it -v <your video file>:/var/input/video -v <output directory>:/var/vod -e FILE_NAME=<file name> --rm hls-ffmpeg`