# liber
*Go library for creating EPUB files*

## Description
- This library provides a high-level, ergonomic API for creating EPUB files (2.0.1). 
- It covers all [epubcheck](https://github.com/w3c/epubcheck) validations

## Installation
```bash
go get -u github.com/javiorfo/go-liber@latest
```

## Example

```go
package main

import (
  "log"
  "os"

  "github.com/javiorfo/go-liber"
  "github.com/javiorfo/go-liber/body"
  "github.com/javiorfo/go-liber/ident"
  "github.com/javiorfo/go-liber/lang"
  "github.com/javiorfo/go-liber/reftype"
  "github.com/javiorfo/go-liber/resource"
)

func main() {
  f, err := os.Create("book.epub")
  if err != nil {
	  panic(err)
  }
  defer f.Close()

  e := liber.EpubBuilder(
	  liber.MetadataBuilder("My Book", lang.Spanish, ident.Default()).
		  Creator("Johan Gambolputty").
		  Build(),
  ).
	  Stylesheet(body.Raw("body {}")).
	  CoverImage(resource.JpgFile("/path/cats.jpg")).
	  AddResources(resource.PngFile("/path/cs.png")).
	  AddContents(
		  liber.ContentBuilder(body.Raw(
          `<h1>Chapter 1</h1>
		  <h2 id="id01">Section 1.1</h2>
		  <h2 id="id02">Section 1.1.1</h2>
		  <h2 id="id03">Section 1.2</h2>`,
          ), reftype.Text("Chapter 1")).
			  AddContentReferences(liber.ContentReferenceBuilder("Section 1.1").
				  AddChildren(liber.ContentReferenceBuilder("Section 1.1.1").Build()).
				  Build()).
			  AddContentReferences(liber.ContentReferenceBuilder("Section 1.2").Build()).
			  AddChildren(
				  liber.ContentBuilder(body.Raw("<h1>Chapter 2</h1>"), reftype.Text("Chapter 2")).Build(),
				  liber.ContentBuilder(body.Raw("<h1>Chapter 3</h1>"), reftype.Text("Chapter 3")).
					  AddChildren(liber.ContentBuilder(body.Raw("<h1>Chapter 4</h1>"), reftype.Text("Chapter 4")).
						  Build()).
					  Build(),
			  ).
			  Build(),
	  ).
	  Build()

  if err := liber.Create(&e, f); err != nil {
      os.Remove("book.epub")
	  log.Fatal(err)
  }
}
```

## Details
- Every content is a **xhtml**. The entire xhml text or only the body could be added as content (the latter is more practical and secure because follows the standard). See [examples](https://github.com/javiorfo/go-liber/tree/master/examples))
- Content (Ex: Chapter) and ContentReference (Ex: Chapter#ref1) could be named with filename and id methods respectively. If none is set, Content will be sequential c{number}.xhtml (c01.xhtml, c02.xhtml...) and ContentReferences will be id{number} (id01, id02...) corresponding to the Content.


---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
