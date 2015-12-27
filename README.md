[ID3v1]: https://en.wikipedia.org/wiki/ID3#ID3v1
[ID3v2]: https://en.wikipedia.org/wiki/ID3#ID3v2
[MP3]: https://en.wikipedia.org/wiki/MP3

# `mp3`
A simple library to extract metadata ([ID3v1][ID3v1] & [ID3v2][ID3v2]) from [MP3][MP3] files. [**Windmill.mp3**](https://github.com/CentaurWarchief/mp3/blob/master/res/Windmill.mp3) audio sample was taken from [http://eng.universal-soundbank.com](http://eng.universal-soundbank.com/air.htm), it is/will be used for testing purposes only.

### Installation
```
go get github.com/CentaurWarchief/mp3
```

### Features
- It supports all known ID3v2 versions: 2.2, 2.3 and 2.4;
- It only allocates memory for frames that you `Read()`;
- It's entirely based upon native `io.Reader`.


##### References
- http://www.codeproject.com/Articles/8295/MPEG-Audio-Frame-Header
- https://en.wikipedia.org/wiki/ID3#Layout
- http://id3.org/Developer%20Information
