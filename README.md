# CD Flow

Simple framework and binary to emit Cloud Events related to Continuous Delivery. 

The idea behind this framework and command-line interface is to empower you  emitting events related to a Continuous Delivery flow to measure performance and gain visibility of your processes. 

## Supported Events

Use `cdf --help`

The following events are currently supported:
- [Issue Events]()
- [Repository Events]()
- [Artifact Events]()
- [Container Events]()
- [Service Events]()
- [Environment Events]()


## CDF Issue Events

Use `cdf issue --help`

- Base
  - Id
  - Title
  - Repository
  - Author
- Lifecycle Events
  - **CDF.Issue.Created**
    Use `cdf issue created --help`
  - **CDF.Issue.Updated**
    Use `cdf issue updated --help`  
  - **CDF.Issue.CommentAdded**
    Use `cdf issue commented --help`
  - **CDF.Issue.Closed**
    Use `cdf issue closed --help`

## CDF Repository Events
- Base
  - Name
  - Url 
- PR Events
  - **CDF.Repository.PR.Created**
    - PR Url
    - PR Title
    - Issue
  - **CDF.Repository.PR.Merged**
  - **CDF.Repository.PR.CommentAdded**
- Main Events
  - **CDF.Repository.Main.Changed**
- Tags Events
  - **CDF.Repository.Tag.Created**

## CDF Artifact Events
- Base
  - Name
  - Version
  - Source
- Lifecycle Events:
  - **CDF.Artifact.TestStarted**
  - **CDF.Artifact.TestEnded**  
    - TestPassed
  - **CDF.Artifact.Built**
    - SHA 
  - **CDF.Artifact.VersionUpdated**
    - OldVersion
    - NewVersion
  - **CDF.Artifact.Released**
    - URL
    - Kind
      - Library
      - Service
## CDF Container Events
- Base
  - Name
  - Organization
  - Tag
  - Repository
- Lifecycle Events
  - **CDF.Container.Built**
  - **CDF.Container.Released**
  
## CDF Service Events
- Base
  - Name
  - Url
  - Version
  - Environment URL
- Lifecycle Events
  - **CDF.Service.Deployed**
  - **CDF.Service.Upgraded**

## CDF Environment Events
- Base
  - URL
  - Repository
  - Name
- Lifecycle Events
  - **CDF.Environment.Created**
  - **CDF.Environment.Updated**
  - **CDF.Environment.ServicePromoted**


## Visualization
(TBD)

