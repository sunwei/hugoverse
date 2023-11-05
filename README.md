[Sun Wei]: https://github.com/sunwei
[hugo]: https://github.com/sunwei/hugoverse
[go]: https://go.dev/
[hugoverse config api demo]: https://youtu.be/8bV5DyZGAfA
[Hugoverse config api message flow]: https://dddplayer.com/?path=https://assets.dddplayer.com/resource/hugoverse/github.com.sunwei.hugoverse.internal.domain.messageflow.dot
[Hugoverse DDD strategic model v0]: https://dddplayer.com/?path=https://assets.dddplayer.com/resource/hugoverse/github.com.sunwei.hugoverse.strategic.dot
[Hugo allconfig message flow]: https://dddplayer.com/?path=https://assets.dddplayer.com/resource/hugo/github.com.gohugoio.hugo.config.allconfig.messageflow.dot
[Hugo allconfig key objects composition relations]: https://dddplayer.com/?path=https://assets.dddplayer.com/resource/hugo/github.com.gohugoio.hugo.config.allconfig.composition.dot

A [hugo] headless CMS built with love by [Sun Wei] in [Go].

---

### Build

```shell
go install
go build -o hugov
```

### Run

```shell
# example: ./hugov server -p ~/github/sunwei/xyz
./hugov server -p /path/to/hugo/project

# open browser and check the api calling result
open http://localhost:1314/config
```

online video : [hugoverse config api demo]

## Onboard

Join us, let's make [hugo] dynamic from today!

* Check [Hugoverse config api message flow] to understand how the api works.
* Check [Hugoverse DDD strategic model v0] to see current domain model
* Check [Hugo allconfig message flow] to understand Hugo source code message flow
* Check [Hugo allconfig key objects composition relations] to understand Hugo source code relations

