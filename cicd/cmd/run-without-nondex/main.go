package main

import (
	"flag"
	"log"

	"github.com/GoogleCloudPlatform/DataflowTemplates/cicd/internal/flags"
	"github.com/GoogleCloudPlatform/DataflowTemplates/cicd/internal/workflows"
)

func main() {
	flags.RegisterCommonFlags()
	flag.Parse()

	mvnFlags := workflows.NewMavenFlags()
	if err := workflows.RunSingleTestCheck().Run(
		mvnFlags.IncludeDependencies(),
		mvnFlags.IncludeDependents(),
		mvnFlags.SkipCheckstyle(),
		mvnFlags.SkipSpotlessCheck(),
		mvnFlags.SpecifyNonDexTest("com.google.cloud.teleport.templates.common.DatastoreConvertersTest#testCheckNoKeyAllInvalid")); err != nil {
		log.Fatalf("Error running test without NonDex check: %v", err)
	}
	log.Println("Test check without NonDex completed successfully!")

}
