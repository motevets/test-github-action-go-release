# springboard

_a server and client implementing Spring '83_

This server and client aim to track the latest drafts of the spring89 protocol specification.
Currently this implementation is targetting [draft-20220616](https://github.com/robinsloan/spring-83-spec/blob/main/draft-20220616.md).

You can see it in action at [spring83.kindrobot.ca](https://spring83.kindrobot.ca). (Come say hi!)

## Post a board

You can download precompiled binaries for your system / architecture from the releases tab.

Run `./springboard help` (or `.\springboard.exe help` on Windows) to get started posting to a spring83 server.

An example:

```bash
./springboard generate-keys
echo "<p>Hello, world!</p>" > board.html
./springboard post https://spring83.kindrobot.ca < board.html
```

`./springboard generate-keys` may take several minutes and use a lot of proccessing power.
By default, it will save the key pair to `$HOME/.config/spring83`. 

`board.html` can be any valid-ish HTML5 document. It may not have scripts or
load externl resources. You should not put:

```html
<time datatime="...">
```

in `board.html`, springboard will do this for you.

### Help, my key doesn't work

The key format changed between draft versions. You may need to upgrade your client to v1 or greater, delete the contents of `~/.config/spring83`, and try again.

## Hacking

### run the server

Run `PORT=8000 go run cmd/s83server/main.go`.  `PORT` is optional and defaults to 8000.

### run the client

On first run, the client will generate a keypair for you according to the spring83 spec, and store it in `~/.config/spring83/key.pub` and `~/.config/spring83/key.priv`.

This key has to meet a certain specification, so it may take some time to generate on the first run.

`echo "<em>very</em> <pre>cool</pre>" | go run cmd/s83client/main.go http://localhost:8000`

Run `go run cmd/s83client/main.go --help` for all options.

### view the content

go to http://localhost:8000 while the server is running

## Other known Spring '83 implementations
| Name                       | Lang                | Instance                 |
| -------------------------- | ------------------- | -------------------------|
| [llimllib/springer]        | golang              |                          |
| [michael-lazar/lets-dance] | python              | https://spring83.mozz.us |
| [rpj/spring83]             | Javascript (NodeJS) | https://0l0.lol/         |

## Credits

- this package was forked from [llimllib/springer](https://github.com/llimllib/springer/)
- the protocol (and original newsletter) is/was by [Robin Sloan](https://www.robinsloan.com/lab/specifying-spring-83/)

## License

This work is released into the public domain.

[llimllib/springer]: https://github.com/llimllib/springer
[michael-lazar/lets-dance]: https://github.com/michael-lazar/lets-dance
[rpj/spring83]: https://github.com/rpj/spring83
