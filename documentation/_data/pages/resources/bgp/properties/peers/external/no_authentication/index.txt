---
page_title: "peers.external.no_authentication"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "peers external no authentication"], "body_bytes": 1092, "body_sha256": "sha256:5b348299c2538a5721910de8d855db9d01f7375d0c7cb14dde2b3f28c5c25921", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:no_authentication", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external", "path": "documentation/resources/bgp/properties/peers/external/no_authentication/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0330223011231313-0131302323100133-2232102110312200-0331033133303032-1301100120202101-2320101223111322-3100000322300223-0032212000310102", "registry_path": "docs/guides/resources--bgp--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "no_authentication"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/no_authentication/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["bgpCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.no_authentication

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- peers.external.no_authentication

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authentication.

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
no_authentication = {}
```

This is an empty object or choice marker. It has no direct properties.
