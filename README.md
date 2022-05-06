cdc - CLI text codec
====================


Quick start
-----------

```shell
$ make cdc        # generate cdc binary
$ make install    # install
```

Features
--------

- url (en|de)code
- url quick (en|de)code
- xml entity (en|de)code


Options
-------

- `-u` - URL format (default)
- `-x` - XML format
- `-e` - encode (default)
- `-d` - decode
- `-a` - all chars
- `-s` - encode-strip (drop everything except alnums and dashes, replace spaces with dashes, lowercased)


Usage
-----

```
$ cdc -ue 'Gol@ng?&' # go's PathEscape
Gol@ng%3F&

$ cdc -uae 'Gol@ng?&' # go's QueryEscape
Gol%40ng%3F%26

$ cdc -de 'Gol@ng%3F&'
Gol@ng?&

$ cdc -dae 'Gol%40ng%3F%26'
Gol@ng?&

$ cdc -s 'Gol@ng?& says wat'
golng-says-wat
```
