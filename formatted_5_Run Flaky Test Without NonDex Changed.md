2024-02-14T16:42:33.2771430Z ##[group]Run ./cicd/run-without-nondex
2024-02-14T16:42:33.2772759Z [36;1m./cicd/run-without-nondex[0m
2024-02-14T16:42:33.2831357Z shell: /usr/bin/bash -e {0}
2024-02-14T16:42:33.2832149Z env:
2024-02-14T16:42:33.2833342Z   MAVEN_OPTS: -Dorg.slf4j.simpleLogger.log.org.apache.maven.plugins.shade=error
2024-02-14T16:42:33.2835385Z   GOROOT: /home/fedja/Desktop/my_runner/actions-runner/_work/_tool/go/1.21.7/x64
2024-02-14T16:42:33.2837679Z   JAVA_HOME: /home/fedja/Desktop/my_runner/actions-runner/_work/_tool/Java_Temurin-Hotspot_jdk/11.0.22-7/x64
2024-02-14T16:42:33.2839436Z   TODAY: 20240214
2024-02-14T16:42:33.2840132Z   YESTERDAY: 20240213
2024-02-14T16:42:33.2840834Z ##[endgroup]
2024-02-14T16:42:33.2966760Z 2024/02/14 17:42:33 Running command: mvn -B test -f pom.xml -e -pl v1 -Dcheckstyle.skip -Dspotless.check.skip -Dtest=com.google.cloud.teleport.templates.common.DatastoreConvertersTest#testCheckNoKeyAllInvalid
2024-02-14T16:42:36.8861393Z [**INFO**] Error stacktraces are turned on.
2024-02-14T16:42:36.9423798Z [**INFO**] Scanning for projects...
2024-02-14T16:42:38.5093276Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T16:42:38.5096326Z [**INFO**] Detecting the operating system and CPU architecture
2024-02-14T16:42:38.5102696Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T16:42:38.5107801Z [**INFO**] os.detected.name: linux
2024-02-14T16:42:38.5110218Z [**INFO**] os.detected.arch: x86_64
2024-02-14T16:42:38.5112631Z [**INFO**] os.detected.bitness: 64
2024-02-14T16:42:38.5115903Z [**INFO**] os.detected.version: 6.5
2024-02-14T16:42:38.5118472Z [**INFO**] os.detected.version.major: 6
2024-02-14T16:42:38.5121341Z [**INFO**] os.detected.version.minor: 5
2024-02-14T16:42:38.5140939Z [**INFO**] os.detected.release: ubuntu
2024-02-14T16:42:38.5143049Z [**INFO**] os.detected.release.version: 22.04
2024-02-14T16:42:38.5145699Z [**INFO**] os.detected.release.like.ubuntu: true
2024-02-14T16:42:38.5148428Z [**INFO**] os.detected.release.like.debian: true
2024-02-14T16:42:38.5150969Z [**INFO**] os.detected.classifier: linux-x86_64
2024-02-14T16:42:38.5221742Z [**INFO**] 
2024-02-14T16:42:38.5225609Z [**INFO**] ------------< com.google.cloud.teleport:classic-templates >-------------
2024-02-14T16:42:38.5227137Z [**INFO**] Building Google Cloud Teleport Classic 1.0-SNAPSHOT
2024-02-14T16:42:38.5230113Z [**INFO**] --------------------------------[ jar ]---------------------------------
2024-02-14T16:42:39.7630540Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:39.7638554Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:40.9818710Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/maven-metadata.xml (789 B at 645 B/s)
2024-02-14T16:42:40.9963457Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.pom
2024-02-14T16:42:41.3842613Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.pom
2024-02-14T16:42:41.6742427Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.pom (3.3 kB at 11 kB/s)
2024-02-14T16:42:41.6812132Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/templates/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:41.6819631Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:41.9459981Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/maven-metadata.xml (609 B at 2.3 kB/s)
2024-02-14T16:42:42.2176671Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/templates-1.0-20240109.172859-10.pom
2024-02-14T16:42:42.6124725Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/templates-1.0-20240109.172859-10.pom (21 kB at 54 kB/s)
2024-02-14T16:42:47.9518463Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:47.9526839Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:48.2379449Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/maven-metadata.xml (793 B at 2.8 kB/s)
2024-02-14T16:42:48.4528026Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.pom
2024-02-14T16:42:49.2699186Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.pom
2024-02-14T16:42:49.5342825Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.pom (6.4 kB at 24 kB/s)
2024-02-14T16:42:49.5405769Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:49.5413078Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:49.8060599Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/maven-metadata.xml (621 B at 2.3 kB/s)
2024-02-14T16:42:50.0488926Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/integration-testing-lib-1.0-20240109.172859-7.pom
2024-02-14T16:42:50.3108689Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/integration-testing-lib-1.0-20240109.172859-7.pom (6.2 kB at 24 kB/s)
2024-02-14T16:42:51.2999505Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:51.3007355Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T16:42:51.5721948Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/maven-metadata.xml (779 B at 2.9 kB/s)
2024-02-14T16:42:51.8827375Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.pom
2024-02-14T16:42:52.2955517Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.pom
2024-02-14T16:42:52.5667523Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.pom (2.6 kB at 9.4 kB/s)
2024-02-14T16:42:54.9967875Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.jar
2024-02-14T16:42:54.9972127Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.jar
2024-02-14T16:42:54.9975608Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.jar
2024-02-14T16:42:55.6520680Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.jar (23 kB at 36 kB/s)
2024-02-14T16:42:55.6533221Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.jar (41 kB at 62 kB/s)
2024-02-14T16:42:55.6655206Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.jar (245 kB at 366 kB/s)
2024-02-14T16:42:55.6881048Z [**INFO**] 
2024-02-14T16:42:55.6889661Z [**INFO**] --- maven-enforcer-plugin:3.0.0-M1:enforce (enforce) @ classic-templates ---
2024-02-14T16:43:00.8436289Z [**WARNING**] Failure to find io.confluent:kafka-avro-serializer:jar:5.3.2 in https://splunk.jfrog.io/splunk/ext-releases-local was cached in the local repository, resolution will not be reattempted until the update interval of splunk-artifactory has elapsed or updates are forced
2024-02-14T16:43:00.8438297Z 
2024-02-14T16:43:00.8438726Z Try downloading the file manually from the project website.
2024-02-14T16:43:00.8439294Z 
2024-02-14T16:43:00.8439494Z Then, install it using the command: 
2024-02-14T16:43:00.8440972Z     mvn install:install-file -DgroupId=io.confluent -DartifactId=kafka-avro-serializer -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file
2024-02-14T16:43:00.8441850Z 
2024-02-14T16:43:00.8442697Z Alternatively, if you host your own repository you can deploy the file there: 
2024-02-14T16:43:00.8444930Z     mvn deploy:deploy-file -DgroupId=io.confluent -DartifactId=kafka-avro-serializer -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file -Durl=[url] -DrepositoryId=[id]
2024-02-14T16:43:00.8446284Z 
2024-02-14T16:43:00.8446464Z Path to dependency: 
2024-02-14T16:43:00.8447250Z 	1) com.google.cloud.teleport:classic-templates:jar:1.0-SNAPSHOT
2024-02-14T16:43:00.8448372Z 	2) org.apache.beam:beam-runners-google-cloud-dataflow-java:jar:2.53.0
2024-02-14T16:43:00.8449373Z 	3) org.apache.beam:beam-sdks-java-io-kafka:jar:2.53.0
2024-02-14T16:43:00.8450237Z 	4) io.confluent:kafka-avro-serializer:jar:5.3.2
2024-02-14T16:43:00.8450717Z 
2024-02-14T16:43:00.8450728Z 
2024-02-14T16:43:00.8451050Z   io.confluent:kafka-avro-serializer:jar:5.3.2
2024-02-14T16:43:00.8451456Z 
2024-02-14T16:43:00.8451662Z from the specified remote repositories:
2024-02-14T16:43:00.8452763Z   splunk-artifactory (https://splunk.jfrog.io/splunk/ext-releases-local, releases=true, snapshots=true),
2024-02-14T16:43:00.8454018Z   ossrh (https://oss.sonatype.org/content/repositories/snapshots, releases=true, snapshots=true),
2024-02-14T16:43:00.8455100Z   central (https://repo.maven.apache.org/maven2, releases=true, snapshots=false)
2024-02-14T16:43:00.8456317Z 
2024-02-14T16:43:00.8541773Z [**WARNING**] Failure to find io.confluent:kafka-schema-registry-client:jar:5.3.2 in https://splunk.jfrog.io/splunk/ext-releases-local was cached in the local repository, resolution will not be reattempted until the update interval of splunk-artifactory has elapsed or updates are forced
2024-02-14T16:43:00.8543731Z 
2024-02-14T16:43:00.8544658Z Try downloading the file manually from the project website.
2024-02-14T16:43:00.8545313Z 
2024-02-14T16:43:00.8545589Z Then, install it using the command: 
2024-02-14T16:43:00.8548085Z     mvn install:install-file -DgroupId=io.confluent -DartifactId=kafka-schema-registry-client -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file
2024-02-14T16:43:00.8549491Z 
2024-02-14T16:43:00.8549998Z Alternatively, if you host your own repository you can deploy the file there: 
2024-02-14T16:43:00.8552319Z     mvn deploy:deploy-file -DgroupId=io.confluent -DartifactId=kafka-schema-registry-client -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file -Durl=[url] -DrepositoryId=[id]
2024-02-14T16:43:00.8554002Z 
2024-02-14T16:43:00.8554240Z Path to dependency: 
2024-02-14T16:43:00.8555236Z 	1) com.google.cloud.teleport:classic-templates:jar:1.0-SNAPSHOT
2024-02-14T16:43:00.8556671Z 	2) org.apache.beam:beam-runners-google-cloud-dataflow-java:jar:2.53.0
2024-02-14T16:43:00.8557703Z 	3) org.apache.beam:beam-sdks-java-io-kafka:jar:2.53.0
2024-02-14T16:43:00.8558443Z 	4) io.confluent:kafka-schema-registry-client:jar:5.3.2
2024-02-14T16:43:00.8558849Z 
2024-02-14T16:43:00.8558857Z 
2024-02-14T16:43:00.8559182Z   io.confluent:kafka-schema-registry-client:jar:5.3.2
2024-02-14T16:43:00.8559578Z 
2024-02-14T16:43:00.8559757Z from the specified remote repositories:
2024-02-14T16:43:00.8560767Z   splunk-artifactory (https://splunk.jfrog.io/splunk/ext-releases-local, releases=true, snapshots=true),
2024-02-14T16:43:00.8562006Z   ossrh (https://oss.sonatype.org/content/repositories/snapshots, releases=true, snapshots=true),
2024-02-14T16:43:00.8563109Z   central (https://repo.maven.apache.org/maven2, releases=true, snapshots=false)
2024-02-14T16:43:00.8563652Z 
2024-02-14T16:43:13.5886675Z [**INFO**] 
2024-02-14T16:43:13.5889346Z [**INFO**] --- maven-enforcer-plugin:3.0.0-M1:enforce (enforce-banned-dependencies) @ classic-templates ---
2024-02-14T16:43:13.7996702Z [**INFO**] 
2024-02-14T16:43:13.7999579Z [**INFO**] --- jacoco-maven-plugin:0.8.8:prepare-agent (default-prepare-agent) @ classic-templates ---
2024-02-14T16:43:14.2306116Z [**INFO**] argLine set to -javaagent:/home/fedja/.m2/repository/org/jacoco/org.jacoco.agent/0.8.8/org.jacoco.agent-0.8.8-runtime.jar=destfile=/home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/jacoco.exec,excludes=/*InformationSchemaScanner.*:/*AutoValue_*
2024-02-14T16:43:14.2310598Z [**INFO**] 
2024-02-14T16:43:14.2312042Z [**INFO**] --- protobuf-maven-plugin:0.6.1:compile (default) @ classic-templates ---
2024-02-14T16:43:15.0573882Z [**INFO**] Compiling 2 proto file(s) to /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/java
2024-02-14T16:43:15.1253699Z [**INFO**] 
2024-02-14T16:43:15.1255753Z [**INFO**] --- protobuf-maven-plugin:0.6.1:compile-custom (default) @ classic-templates ---
2024-02-14T16:43:15.2871917Z [**INFO**] Compiling 2 proto file(s) to /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/grpc-java
2024-02-14T16:43:15.3043948Z [**INFO**] 
2024-02-14T16:43:15.3045813Z [**INFO**] --- avro-maven-plugin:1.11.3:schema (default) @ classic-templates ---
2024-02-14T16:43:15.7654061Z [**WARNING**] Ignored the com.google.cloud.teleport.bigtable.BigtableCell.timestamp.logicalType property ("timestamp-micros"). It should probably be nested inside the "type" for the field.
2024-02-14T16:43:16.0227998Z [**INFO**] 
2024-02-14T16:43:16.0231212Z [**INFO**] --- maven-resources-plugin:2.6:resources (default-resources) @ classic-templates ---
2024-02-14T16:43:16.3428952Z [**INFO**] Using 'UTF-8' encoding to copy filtered resources.
2024-02-14T16:43:16.3459924Z [**INFO**] Copying 5 resources
2024-02-14T16:43:16.3664245Z [**INFO**] Copying 2 resources
2024-02-14T16:43:16.3672343Z [**INFO**] Copying 2 resources
2024-02-14T16:43:16.3676694Z [**INFO**] 
2024-02-14T16:43:16.3682024Z [**INFO**] --- maven-compiler-plugin:3.11.0:compile (default-compile) @ classic-templates ---
2024-02-14T16:43:16.9384291Z [**INFO**] Changes detected - recompiling the module! :dependency
2024-02-14T16:43:16.9413468Z [**INFO**] Compiling 164 source files with javac [debug target 11] to target/classes
2024-02-14T16:43:24.0070476Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/main/java/com/google/cloud/teleport/splunk/HttpEventPublisher.java:[108,19] [AutoValueMutable] An @AutoValue property that is a primitive array returns the original array, which can therefore be modified by the caller. If this is OK, you can suppress this warning with @SuppressWarnings("mutable"). Otherwise, you should replace the property with an immutable type, perhaps a simple wrapper around the original array.
2024-02-14T16:43:24.0077297Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/java/com/google/cloud/teleport/spanner/proto/ExportProtos.java: Some input files use or override a deprecated API.
2024-02-14T16:43:24.0081817Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/java/com/google/cloud/teleport/spanner/proto/ExportProtos.java: Recompile with -Xlint:deprecation for details.
2024-02-14T16:43:24.0086196Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/main/java/com/google/cloud/teleport/templates/DatastoreToText.java: Some input files use unchecked or unsafe operations.
2024-02-14T16:43:24.0090237Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/main/java/com/google/cloud/teleport/templates/DatastoreToText.java: Recompile with -Xlint:unchecked for details.
2024-02-14T16:43:24.0092333Z [**INFO**] 
2024-02-14T16:43:24.0093372Z [**INFO**] --- maven-resources-plugin:2.6:testResources (default-testResources) @ classic-templates ---
2024-02-14T16:43:24.0108921Z [**INFO**] Using 'UTF-8' encoding to copy filtered resources.
2024-02-14T16:43:24.0134010Z [**INFO**] Copying 34 resources
2024-02-14T16:43:24.0416309Z [**INFO**] 
2024-02-14T16:43:24.0418539Z [**INFO**] --- maven-compiler-plugin:3.11.0:testCompile (default-testCompile) @ classic-templates ---
2024-02-14T16:43:24.0527658Z [**INFO**] Changes detected - recompiling the module! :dependency
2024-02-14T16:43:24.0537788Z [**INFO**] Compiling 136 source files with javac [debug target 11] to target/test-classes
2024-02-14T16:43:27.8917850Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/spanner/ddl/RandomDdlGenerator.java:[107,25] [AutoValueMutable] An @AutoValue property that is a primitive array returns the original array, which can therefore be modified by the caller. If this is OK, you can suppress this warning with @SuppressWarnings("mutable"). Otherwise, you should replace the property with an immutable type, perhaps a simple wrapper around the original array.
2024-02-14T16:43:27.8925339Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/templates/SpannerToTextIT.java:[114,57] non-varargs call of varargs method with inexact argument type for last parameter;
2024-02-14T16:43:27.8928084Z   cast to java.lang.Object for a varargs call
2024-02-14T16:43:27.8929261Z   cast to java.lang.Object[] for a non-varargs call and to suppress this warning
2024-02-14T16:43:27.8932532Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/templates/SpannerToTextIT.java:[173,57] non-varargs call of varargs method with inexact argument type for last parameter;
2024-02-14T16:43:27.8935884Z   cast to java.lang.Object for a varargs call
2024-02-14T16:43:27.8937133Z   cast to java.lang.Object[] for a non-varargs call and to suppress this warning
2024-02-14T16:43:27.8940455Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/util/DatadogApiKeyNestedValueProviderTest.java: Some input files use or override a deprecated API.
2024-02-14T16:43:27.8945732Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/util/DatadogApiKeyNestedValueProviderTest.java: Recompile with -Xlint:deprecation for details.
2024-02-14T16:43:27.8950595Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/spanner/ExportTransformTest.java: Some input files use unchecked or unsafe operations.
2024-02-14T16:43:27.8955276Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/spanner/ExportTransformTest.java: Recompile with -Xlint:unchecked for details.
2024-02-14T16:43:27.8957505Z [**INFO**] 
2024-02-14T16:43:27.8958426Z [**INFO**] --- maven-checkstyle-plugin:3.2.1:check (default) @ classic-templates ---
2024-02-14T16:43:28.8968555Z [**INFO**] 
2024-02-14T16:43:28.8969732Z [**INFO**] --- maven-surefire-plugin:2.21.0:test (default-test) @ classic-templates ---
2024-02-14T16:43:29.1835197Z [**INFO**] Surefire report directory: /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/surefire-reports
2024-02-14T16:43:29.1861956Z [**INFO**] parallel='none', perCoreThreadCount=true, threadCount=0, useUnlimitedThreads=false, threadCountSuites=0, threadCountClasses=0, threadCountMethods=0, parallelOptimized=true
2024-02-14T16:43:29.3726356Z [**INFO**] 
2024-02-14T16:43:29.3727313Z [**INFO**] -------------------------------------------------------
2024-02-14T16:43:29.3728006Z [**INFO**]  T E S T S
2024-02-14T16:43:29.3728832Z [**INFO**] -------------------------------------------------------
2024-02-14T16:43:30.8015466Z [**INFO**] Running com.google.cloud.teleport.templates.common.DatastoreConvertersTest
2024-02-14T16:43:30.8188240Z SLF4J: Class path contains multiple SLF4J bindings.
2024-02-14T16:43:30.8194073Z SLF4J: Found binding in [jar:file:/home/fedja/.m2/repository/org/slf4j/slf4j-reload4j/1.7.36/slf4j-reload4j-1.7.36.jar!/org/slf4j/impl/StaticLoggerBinder.class]
2024-02-14T16:43:30.8199440Z SLF4J: Found binding in [jar:file:/home/fedja/.m2/repository/org/slf4j/slf4j-log4j12/1.7.25/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
2024-02-14T16:43:30.8204352Z SLF4J: Found binding in [jar:file:/home/fedja/.m2/repository/org/slf4j/slf4j-simple/1.7.36/slf4j-simple-1.7.36.jar!/org/slf4j/impl/StaticLoggerBinder.class]
2024-02-14T16:43:30.8207795Z SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
2024-02-14T16:43:30.9935485Z SLF4J: Actual binding is of type [org.slf4j.impl.Reload4jLoggerFactory]
2024-02-14T16:43:36.2707833Z [**INFO**] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 5.448 s - in com.google.cloud.teleport.templates.common.DatastoreConvertersTest
2024-02-14T16:43:36.7556122Z [**INFO**] 
2024-02-14T16:43:36.7557112Z [**INFO**] Results:
2024-02-14T16:43:36.7558073Z [**INFO**] 
2024-02-14T16:43:36.7559326Z [**INFO**] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0
2024-02-14T16:43:36.7560841Z [**INFO**] 
2024-02-14T16:43:36.7616199Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T16:43:36.7617766Z [**INFO**] BUILD SUCCESS
2024-02-14T16:43:36.7619582Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T16:43:36.7620897Z [**INFO**] Total time:  59.846 s
2024-02-14T16:43:36.7623371Z [**INFO**] Finished at: 2024-02-14T17:43:36+01:00
2024-02-14T16:43:36.7625571Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T16:43:36.9017124Z 2024/02/14 17:43:36 Test check without NonDex completed successfully!
