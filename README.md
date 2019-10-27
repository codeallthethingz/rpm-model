# RPM Model

Version: 0.2.0

A set of models that the RPM framework uses to think about 3D reality.

The file `model/construct-json.go` describes objects stored in the central repository.

An autogenerated JS version is created in the `model-js/` folder and published to npm
 
See the `model/examples` for how this format fits together.

## Development

Install task

`brew install go-task/tap/go-task`

Run bump-patch

`task bump-patch`

Publish to npm

`task publish`
