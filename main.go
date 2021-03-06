package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	skipInstallation bool
)

func main() {
	flag.BoolVar(&skipInstallation, "skip-install", false, "Specify whether or not you want to skip executing 'npm install'")
	flag.Parse()
	g := &Generator{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("--- Welcome to use this Flux Generator ---\n\n")
	fmt.Println("What is your App Name ?")
	appName, _ := reader.ReadString('\n')

	appName = appName[:len(appName)-1]
	g.SetName(appName)

	fmt.Println("What is the author of this app ?")
	authorName, _ := reader.ReadString('\n')
	authorName = authorName[:len(authorName)-1]
	g.SetAuthorName(authorName)

	g.CreateDirectories(
		"Generating folder structure now....",

		[]string{
			"app/",
			"app/",
			"app/components/",
			"app/actions/",
			"app/containers/",
			"app/reducers/",
			"app/constants/",
			"app/middlewares/",
			"app/test/",
			"app/test/actions/",
			"app/test/components/",
			"app/test/reducers/",
			"static/",
		})

	g.WriteFiles(
		"Generating necessary files now....",

		[][2]string{
			//			[2]string{"src/scripts/dispatcher/AppDispatcher.js", dispatcherFileContent},
			//			[2]string{"src/scripts/components/App.js", appJSFileContent},
			[2]string{"index.js", mainJSFileContent},
			[2]string{"webpack.config.js", webpackConfigFileContent},
			[2]string{"package.json", packageJSONFileContent(g.GetName(), g.GetAuthorName())},
			[2]string{"index.html", indexHTMLFileContent},
			[2]string{"server.js", serverJsFileContent},
			[2]string{"app/containers/Root.js", containerRootContent},
			[2]string{"app/containers/App.js", containerAppContent},
			[2]string{"app/reducers/index.js", reducerIndexContent},
			[2]string{"app/reducers/todos.js", reducerTodosContent},
		})

	var changeDirError error
	if skipInstallation == false {
		changeDirError = os.Chdir(appName)

		if changeDirError == nil {
			g.Run(
				"Executing 'npm install' now....",
				"npm", "install",
			)
			g.Run(
				"Executing 'webpack' now....",
				"webpack",
			)
		}
	}

	if g.Err() != nil {
		// If there were any error occured when doing one of the tasks above, all the tasks behind this task will not be executed.
		// You should handle this error here only once.
		fmt.Println(g.Err())
	} else if changeDirError != nil {
		fmt.Println(changeDirError)
	} else {
		fmt.Println("All the tasks have been done sucessfully.")
	}
}
