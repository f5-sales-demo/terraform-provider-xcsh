---
page_title: "private_connectivity.cloud_link"
subcategory: "Infrastructure"
description: "Information related to cloud link used by the site."
xcsh_docs: {"aliases": ["private connectivity cloud link"], "body_bytes": 1577, "body_sha256": "sha256:f6c138f238382aef9ac2256c2c3f85432d81a3c75339c81f9433afbb5b011785", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "parent_id": "xcsh-docs:data-sources:site:properties:private_connectivity", "path": "documentation/data-sources/site/properties/private_connectivity/cloud_link/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3322103333002101-3033211000122021-0021323121112201-3033133313133102-1112112302221031-1310113231122321-0113200232320010-3113000123300120", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["private_connectivity", "cloud_link"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-private_connectivity--cloud_link--name", "description": "Name of the the CloudLink used with this site.", "document_id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "cloud_link", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["state"], "anchor": "schema-private_connectivity--cloud_link--state", "description": "State of the CloudLink connections - UP: Up CloudLink and their corresponding Direct Connect connections are up and healthy - DOWN: Down CloudLink and their corresponding Direct Connect connections are down - DEGRADED: Degraded Some of Direct Connect connections with the CloudLink are down .. Possible values are `UP`,", "document_id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "cloud_link", "state"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/private_connectivity/cloud_link/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Information related to cloud link used by the site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity.cloud_link

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/private_connectivity/)
- private_connectivity.cloud_link

<a id="section"></a>

Type: `"single"`. Computed.

Information related to cloud link used by the site.

## Direct properties

<a id="schema-private_connectivity--cloud_link--name"></a>

### name property

Type: `"string"`. Computed.

Name of the the CloudLink used with this site.

<a id="schema-private_connectivity--cloud_link--state"></a>

### state property

Type: `"string"`. Computed.

\[Enum: UP|DOWN|DEGRADED|NOT\_APPLICABLE\] State of the CloudLink connections - UP: Up CloudLink and
their corresponding Direct Connect connections are up and healthy - DOWN: Down CloudLink and their
corresponding Direct Connect connections are down - DEGRADED: Degraded Some of Direct Connect
connections with the CloudLink are down .. Possible values are \`UP\`, \`DOWN\`, \`DEGRADED\`,
\`NOT\_APPLICABLE\`. Defaults to \`UP\`.

## Next pages

- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/private_connectivity/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
