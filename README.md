# CD Flow

Simple framework and binary to emit Cloud Events related to Continuous Delivery. 

The idea behind this framework and command-line interface is to empower you  emitting events related to a Continuous Delivery flow to measure performance and gain visibility of your processes. 

## Supported Events

Use `cdf --help`

The following events are currently supported:
- [Repository Events]()
  - [Issue Events]()
  - [Pull Request Events]()
  - [Branch Events]()
  - [Tag Events]()
- [Pipeline Events]()  
  - [Artifact Events]()
  - [Container Events]()
- [Environment Events]()
  - [Service Events]()
- [Infrastructure Events]()  



# CDF Repository Events

Use `cdf repo --help`

- Base Properties
  - Name
  - Url 

### **CDF.Repository.Created** Event 

Example: `./cdf repo created --name my-project --url http://github.com/salaboy/my-project`

### **CDF.Repository.Deleted** Event 

Example: `./cdf repo deleted --name my-project --url http://github.com/salaboy/my-project`

## CDF Issue Events

Use `cdf issue --help`

- Base Properties
  - Id (required)
  - Repository (required)
  - Title 
  - Author 
  
### **CDF.Issue.Created** Event

Example: `./cdf issue created --id 7 --repository my-project --title "new issue" --author salaboy`

### **CDF.Issue.Updated** Event

Example: `./cdf issue updated --id 7 --repository my-project`

### **CDF.Issue.CommentAdded** Event

Example: `./cdf issue commented --id 7 --repository my-project --comment "new comment"`

### **CDF.Issue.Closed** Event

Example: `./cdf issue closed --id 7 --repository my-project`


## Pull Request Events

Use `cdf pr --help`

- Base Properties    
    - PR Id (required)
    - PR Repository (required)
    - Issue Id
    - PR Title


### **CDF.PR.Created** Event

Example: `./cdf pr created --id 42 --repository my-project --author salaboy --title "fixing issue 7"`

### **CDF.PR.Merged** Event

Example: `./cdf pr merged --id 7 --repository my-project`

### **CDF.PR.CommentAdded** Event

Example: `./cdf pr commented --id 7 --repository my-project --comment "hi there from a pr comment"`


## CDF Branch Events

Use `cdf branch --help`

- Base Properties
  - Repository
  - Name
  - Url 

### **CDF.Branch.Created** Event 

Example: `./cdf branch created --name my-branch --repository my-project`

### **CDF.Branch.Deleted** Event 

Example: `./cdf branch deleted --name my-branch --repository my-project`

## CDF Tag Events

Use `cdf tag --help`

- Base Properties
  - Repository
  - Name
  - Url 

### **CDF.Tag.Created** Event 

Example: `./cdf tag created --name 0.0.1 --repository my-project`

### **CDF.Tag.Deleted** Event 

Example: `./cdf tag deleted --name 0.0.1 --repository my-project`


## Pipeline Events

Use `cdf pipeline --help`

- Base Properties
  - Id
  - Name
  - Repository

### **CDF.Pipeline.Started** Event

Example: `./cdf pipeline started --name my-service-pipeline --id UUID-abc-123 --repository my-project`

### **CDF.Pipeline.Finished** Event

Example: `./cdf pipeline finished --name my-service-pipeline --id <UUID> --repository my-project`

### **CDF.Pipeline.Failed** Event

Example: `./cdf pipeline failed --name my-service-pipeline --id <UUID> --repository my-project`


## Artifact Events

Use `cdf artifact --help`

- Base Properties
  - Name
  - Version
  - Source

### **CDF.Artifact.TestStarted** Event

### **CDF.Artifact.TestEnded** Event 
    - Result

### **CDF.Artifact.Built** Event 
    - SHA 
### **CDF.Artifact.VersionUpdated** Event
    - OldVersion
    - NewVersion
### **CDF.Artifact.Released** Event
    - URL
    - Kind
      - Library
      - Service
      
      
## Container Events
- Base Properties
  - Name
  - Organization
  - Tag
  - Repository

### **CDF.Container.Built** Event

### **CDF.Container.Released** Event
  
## Service Events
- Base Properties
  - Name
  - Url
  - Version
  - Environment URL

### **CDF.Service.Deployed** Event

### **CDF.Service.Upgraded** Event

## Environment Events
- Base Properties
  - URL
  - Repository
  - Name

### **CDF.Environment.Created**

### **CDF.Environment.Updated**

### **CDF.Environment.ServicePromoted**

# Metrics
(TBD)

# Visualization

(TBD)

