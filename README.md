# on

## Abount

it's a very simple project to upload of files for aws s3

## Use

The purpose is to use the project to make file uploads to s3 based website

## Flags

- title
- bucket
- folder

### Simple upload

```
    on post.html -b mybucket.com 
```

### Specifying folder
```
    on post.html -b mybucket.com -f blog
```

Note: if name is already exist if name is already exist the file will be replaced

### Specifying title
```
    on post.html -b mybucket.com -f blog -t how-to-upload-file.html
```

Note: saved in s3 as blog/how-to-upload-file.html
