# Pixawttr

![example image of a train in a snowy landscape with a composite addition of current conditions in Golden, CO, USA at a random time](assets/main_img.png)

Pixawttr is a golang program to automate the creation of a pretty composite forcast image.

It uses the [wttr.in](https://wttr.in) service to get "Current conditions" data and an analogous PNG.

It uses the [Pixabay Search API](https://pixabay.com/api/docs/#api_search_images) to enable the random downloading of an image based on a query string.

## Installation

```bash
go get github.com/erindatkinson/pixawttr
```

### Prerequisites

Pixawttr has an external dependency of imagemagick.

#### MacOS/Linux(linuxbrew)

```bash
brew bundle
```

#### Windows

1. Follow the [instructions](https://imagemagick.org/script/download.php#windows)
1. Make sure that the tools are accessible and set in the PATH
  1. If `composite --help` doesn't work in Powershell and the PATH is set correctly, you will also need to Set-Alias for `magick composite` to `composite`.

## Usage

To use pixawttr, you'll need to sign up for a Pixabay account, once you've logged in, you can see your API key in the `key` parameter under the [Search Images](https://pixabay.com/api/docs/#api_search_images) section of the API docs. (I haven't been able to find a way besides this to get the key :woman-shrugging:)

Once you have the key, you can place it in the environment variable `PixabayAPIKey` or at $HOME/.pixawttr in the following format

```text
PixabayAPIKey: 12345678-f4f79534645cd471ea614adf92efffbc
```

after which, you can run

```bash
pixawttr "CityName" [outFile.png]
```

and the resulting image will be either what you've entered for the outFile path, or `outFile.png` in the directory you run it.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

All contributors are held to the [Code of Conduct](CODE_OF_CONDUCT.md) adapted from the [Contributor Covenant version 1.4](https://www.contributor-covenant.org/version/1/4/code-of-conduct.html)

## License
By using this you agree to abide by the [Pixabay License](https://pixabay.com/service/license/) in the usage of images downloaded from Pixabay.

The code for pixawttr is licensed under MPL V2.0 [License](LICENSE)