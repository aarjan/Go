package main

func main() {

	// scrape.Attr(node *html.Node, key string)
	// s := `<p>This is a example html for  parsing<ul><li><a href="/foo"/></li><li><a href="/bar" /></li></ul></p>`
	a := `<html>
		<head/><div id = "a">asdf
			<span class="test">Fist </span>
			<div class="test">Fist </div>
			<div class="test">asd
				<div>Fist</div>
				<div>aaa</div>
			</div>
			<div class="test">fdfd
				<div>abd
					<div>hj</div>
				</div>
				<span>Second</div>
				
			</div>
			<div class="test">Fist </div>
			<div class="test">Fist </div>
			<div class="test">Third</div>
		</div>
		</html>
	// `
	// reader := strings.NewReader(a)
	// doc, err := html.Parse(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var f func(*html.Node)
	// f = func(n *html.Node) {
	// 	if n.Type == html.ElementNode && n.Data == "div" {
	// 		if n.Parent != nil && n.Parent.Parent != nil {
	// 			if scrape.Attr(n.Parent.Parent, "class") == "test" {
	// 				fmt.Println(n.FirstChild.Data)
	// 			}
	// 		}

	// 	}

	// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 		f(c)
	// 	}
	// }

	// f(doc)

}
