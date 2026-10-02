---
page_title: "virtual_server.port_translation"
subcategory: ""
description: "Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any service"
xcsh_docs: {"aliases": ["virtual server port translation"], "body_bytes": 3043, "body_sha256": "sha256:7fabf381e3557d6e5decf96c2a2dd170af06f2d9cbc1d1ea6126e48fc6324e63", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_enable"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/port_translation/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030", "registry_path": "docs/guides/resources--application_profiles--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.port_translation:ConflictingObjectAttributes:port_translation_disable,port_translation_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.port_translation:ConflictingObjectAttributes:port_translation_disable,port_translation_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "port_translation"], "schema_version": 1, "sections": [{"aliases": ["port translation disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "port_translation", "port_translation_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["port translation enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "port_translation", "port_translation_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/port_translation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any service", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.port_translation

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.port_translation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance..

Upstream description:

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port_translation_disable",
    "port_translation_enable")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

Terraform syntax:

```terraform
port_translation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [port_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/): complete subsection reference.

- [port_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/port_translation_enable/): complete subsection reference.

## Next pages

- [virtual_server.port_translation.port_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/)
- [virtual_server.port_translation.port_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/port_translation_enable/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
