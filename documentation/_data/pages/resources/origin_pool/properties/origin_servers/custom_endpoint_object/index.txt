---
page_title: "origin_servers.custom_endpoint_object"
subcategory: "Load Balancing"
description: "origin_servers.custom_endpoint_object for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1588, "body_sha256": "sha256:18adbc32683aa09f73982d82e40e9356d9aa13f894622742763662b92af584a3", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object:endpoint"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "documentation/resources/origin_pool/properties/origin_servers/custom_endpoint_object/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["origin_servers", "custom_endpoint_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/custom_endpoint_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.custom_endpoint_object for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.custom_endpoint_object

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- origin_servers.custom_endpoint_object

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with a reference to endpoint object.

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
custom_endpoint_object {
  # Configure direct properties listed below.
}
```

## Direct properties

- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/endpoint/): complete subsection reference.

## Next pages

- [origin_servers.custom_endpoint_object.endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/endpoint/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
