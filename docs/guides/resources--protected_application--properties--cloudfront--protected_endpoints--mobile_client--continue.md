---
page_title: "cloudfront.protected_endpoints.mobile_client.continue"
subcategory: ""
description: "cloudfront.protected_endpoints.mobile_client.continue for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2274, "body_sha256": "sha256:b0ee03086feb3c1d685c78aaaf97fd8986a85d96949da446436406563470edb8", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:add_header", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:no_header"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.mobile_client.continue for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudfront.protected_endpoints.mobile_client.continue

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client.md)
- cloudfront.protected_endpoints.mobile_client.continue

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_header](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue--add_header.md): complete subsection reference.

- [no_header](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue--no_header.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue--add_header.md)
- [cloudfront.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue--no_header.md)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client.md)
- [xcsh_protected_application](../resources/protected_application.md)
