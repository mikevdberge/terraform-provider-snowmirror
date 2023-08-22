---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "snowmirror Provider"
subcategory: ""
description: |-
  SnowMirror API: This document is a guide which will walk you through SnowMirror REST API.
  The purpose of the REST API is to allow developers to integrate SnowMirror with other applications.
  Manage Synchonization
  Create a synchronizationUpdate a synchronizationRemove a synchronizationExport a synchronizationImport a synchronization
  The following resource collections are offered by this API:
  * Synchonization
  * SynchronizationAction
  * SynchronizationExport
  * SynchronizationImport
---

# snowmirror Provider

SnowMirror API: This document is a guide which will walk you through SnowMirror REST API. 
The purpose of the REST API is to allow developers to integrate SnowMirror with other applications.

__Manage Synchonization__

* Create a synchronization
* Update a synchronization
* Remove a synchronization
* Export a synchronization
* Import a synchronization

The following resource collections are offered by this API:
* Synchonization
* SynchronizationAction
* SynchronizationExport
* SynchronizationImport

## Example Usage

```terraform
terraform {
  required_providers {
    snowmirror = {
      source  = "mikevdberge/snowmirror"
      version = "0.1.1"
    }
  }
}

provider "snowmirror" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `password` (String, Sensitive)
- `server_url` (String) Server URL (defaults to http://localhost:9090)
- `username` (String, Sensitive)