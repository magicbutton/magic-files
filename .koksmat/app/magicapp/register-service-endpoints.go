/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
package magicapp

import (
	"github.com/magicbutton/magic-files/services"
	"github.com/nats-io/nats.go/micro"
)

func RegisterServiceEndpoints(root micro.Group) {
    root.AddEndpoint("user", micro.HandlerFunc(services.HandleUserRequests))
        root.AddEndpoint("platformcategory", micro.HandlerFunc(services.HandlePlatformCategoryRequests))
        root.AddEndpoint("platform", micro.HandlerFunc(services.HandlePlatformRequests))
        root.AddEndpoint("container", micro.HandlerFunc(services.HandleContainerRequests))
        root.AddEndpoint("relation", micro.HandlerFunc(services.HandleRelationRequests))
        root.AddEndpoint("file", micro.HandlerFunc(services.HandleFileRequests))
        root.AddEndpoint("importdata", micro.HandlerFunc(services.HandleImportDataRequests))
        root.AddEndpoint("importsharepointsite", micro.HandlerFunc(services.HandleImportSharePointSiteRequests))
        root.AddEndpoint("importgroups", micro.HandlerFunc(services.HandleImportGroupsRequests))
        root.AddEndpoint("importrelations", micro.HandlerFunc(services.HandleImportRelationsRequests))
        root.AddEndpoint("importsharedmailbox", micro.HandlerFunc(services.HandleImportSharedMailboxRequests))
        root.AddEndpoint("job", micro.HandlerFunc(services.HandleJobRequests))
        root.AddEndpoint("batch", micro.HandlerFunc(services.HandleBatchRequests))
    }
