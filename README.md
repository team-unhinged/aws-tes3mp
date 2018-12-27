# aws-tes3mp

Morrowind TES3MP Server in AWS

Players will need to own a valid copy of Morrowind and set up the lastest release of [tes3mp](https://github.com/TES3MP/openmw-tes3mp/releases)

A copy of the game is not actually needed to run the server.

# Deploy

Make sure you have docker installed.

Run:
- `docker build -t aws-tes3mp .`

```
docker run -it \
	-v ${PWD}/src:/src \
	-e AWS_ACCESS_KEY_ID='your aws secret key id' \
	-e AWS_SECRET_ACCESS_KEY='your aws secret access key' \
	-e AWS_DEFAULT_REGION='your aws preferred region' \
	aws-tes3mp aws cloudformation deploy \
	--template-file src/template.yaml \
	--stack-name aws-tes3mp 
```

*OR for PowerShell users*

```
docker run -it ` 
	-v ${PWD}/src:/src `
	-e AWS_ACCESS_KEY_ID='your aws secret key id' `
	-e AWS_SECRET_ACCESS_KEY='your aws secret access key' `
	-e AWS_DEFAULT_REGION='your aws preferred region' `
	aws-tes3mp aws cloudformation deploy `
	--template-file src/template.yaml `
	--stack-name aws-tes3mp
```

This will deploy a CloudFormation stack into your AWS account.

# Configuration

Currently the server config is hardcoded into the `src/tamplate.yaml` file. You may set the server password, name, maximum players, etc in this file. In the future this will be broken out into a separate config file.

### Nice to have but not here yet

- Server admin cli tool
- IP whitelisting lambda
- Server monitoring lambda - shuts down ec2 when no players for 10min
- Server start up lambda - request instance to be started
