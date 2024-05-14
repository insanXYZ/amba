![video](./__example_video.mp4)

Amba, A Terminal Tool Built with Go for Converting Image Formats (PNG, JPG, JPEG, WEBP)

# Usage
```sh
amba --file <filename> --to <format(use comma(,) if you want convert to many format)>
```

# Requirements

You should install ``libwebp`` on your pc

#### MacOs:
```sh
brew instal webp
```

#### Linux(ubuntu,etc):
```sh
sudo apt-get update
sudo apt-get install libwebp-dev
```

# Supported format
- png
- jpg
- jpeg
- webp