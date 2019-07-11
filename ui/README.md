# Contiv VPP UI

Demo: [video at contivpp.io](https://contivpp.io/demo/contivpp-io-demo-ui/)

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 7.0.2.

## Quickstart

Run your browser with disabled web security (Important for accessing the APIs)

Google Chrome

-Command in OSX:
Open -n -a Google\ Chrome --args --disable-web-security --user-data-dir=/tmp/chrome

-Command in Windows:
In the "Run" app, enter: chrome.exe --user-data-dir="C://Chrome dev session" --disable-web-security

UI can be accessed from web browser at `http://localhost:4300`.

## Configuration

Default UI deployment assumes that Contiv API is exposed via HTTP without additional security features.
If that is not the case follow [README](../k8s/contiv-vpp-ui/README.md) and generate your customized
deployment yaml file using helm.
