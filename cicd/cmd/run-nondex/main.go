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
	if err := workflows.NondexCheck().Run(
		mvnFlags.NonDexNumberOfRuns(10),
		mvnFlags.SpecifyNonDexTest("com.google.cloud.teleport.templates.common.DatastoreConvertersTest#testCheckNoKeyAllInvalid"),
		mvnFlags.SkipCheckstyle(),
		mvnFlags.SkipSpotlessCheck()); err != nil {
		log.Fatalf("Error running NonDex check: %v", err)
	}
	log.Println("NonDex check completed successfully!")

}
