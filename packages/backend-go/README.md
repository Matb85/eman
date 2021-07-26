# Backend server for https://redinnlabs.pl

## required env variables

- PORT - port of the web server - number
- DB_PORT - localhost port of the db - number

## additional env variables for testing

- TEST_DB_BIN_PATH - path to mongod binaries (command for starting mongodb) - string

## commands

### install dependencies

```shell
npm install

npm run bootstrap
```

### build

```
npm run build
```

### start

```
npm start
```
