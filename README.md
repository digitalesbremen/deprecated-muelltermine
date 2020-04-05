# Mülltermine

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://travis-ci.org/digitalesbremen/muelltermine.svg?branch=master)](https://travis-ci.org/digitalesbremen/muelltermine)

Ein in [GO](https://golang.org/) entwickeltes Programm, welches die vom 
[Müllsammler](https://github.com/digitalesbremen/muellsammler) bereitgestellten Abholdaten aller
Bremer Adressen einliest und eine API sowie eine UI bereitsstellt.

Die genauen Hintergründe für diese Anwendung ist auf [Digitales Bremen](https://digitalesbremen.github.io/) beschrieben. 
Für diese Anwendung gibt es eine genauere Beschreibung im [Blog](https://github.com/digitalesbremen/digitalesbremen.github.io/blob/master/blog.md).

## Anforderungen

* Golang 1.14
* make

## Projekt bauen und starten

Das Projekt benötigt eine aktuelle Version von [GO](https://golang.org/) und make. 

```ssh
$ git clone https://github.com/digitalesbremen/muelltermine # Auschecken des Repositories
$ make all                                                  # Bauen einer auf dem Hostsystem laufenden Anwendung
$ ./muellsammler                                            # Starten der Anwendung
```

### Weitere make goals

```
$ make build            # Baut eine auf dem Hostsystem lauffähige Anwendung
$ make test             # Führt die Tests aus
$ make clean            # Räumt all von Go erzeugten Dinge und ggf. existierende Kompilate auf
$ make deps             # Lädt die für die Anwendung notwendigen Go-Abhängigkeiten
$ make build-amd64      # Baut eine amd64 Anwendung (z.B. für OSX)
$ make build-arm        # Baut eine arm Anwendung (z.B. für Rasperry Pi)
```