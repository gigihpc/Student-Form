# babel-plugin-add-filehash

This Babel plugin makes sure the Babel output changes whenever the input
file changes, even if it are only trivial/whitespace changes.

This makes sure subsequent tools like webpack reprocess the file and
show updated eslint messages/...  It specificall works around a webpack
bug that loses error messages on "unchanged" files: https://github.com/webpack/webpack/issues/2538

## Installation

With npm:
```shell
npm install --save-dev babel-plugin-add-filehash
```

Or with yarn:
```shell
yarn --dev babel-plugin-add-filehash
```

After installation of the package you need to add the plugin to
your project's `.babelrc`:

```json
{
  "plugins": ["add-filehash"]
}
```

For more general info about using Babel plugins see the [Babel handbook][1].


 [1]: https://github.com/thejameskyle/babel-handbook/blob/master/translations/en/user-handbook.md#manually-specifying-plugins

