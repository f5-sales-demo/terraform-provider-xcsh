---
page_title: "virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["virtual server immediate action on service down immediate action on service down reset"], "body_bytes": 1362, "body_sha256": "sha256:5851386f6c4ee131a3df7c5eee7219533fc9edcc3afaed1fcbac84fbcaa1a63d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "path": "documentation/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_reset/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1020323002022302-2031200012130221-0130123012221301-0303002023310122-1011303313130223-0213220032010232-0220130012302202-3020330121013201", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_reset"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_reset/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["application_profilesCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.immediate_action_on_service_down](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
immediate_action_on_service_down_reset = {}
```

This is an empty object or choice marker. It has no direct properties.
