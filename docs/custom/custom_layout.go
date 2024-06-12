package custom

import "fmt"

var CustomLayoutJS = fmt.Sprintf(`
	// dark mode
	const style = document.createElement('style');
	style.innerHTML = %s;
	document.head.appendChild(style);
  `, "`"+customCSS+"`")
