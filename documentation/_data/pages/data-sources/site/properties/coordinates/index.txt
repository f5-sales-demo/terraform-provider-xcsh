---
page_title: "coordinates"
subcategory: "Infrastructure"
description: "Coordinates of the site which provides the site physical location."
xcsh_docs: {"aliases": ["coordinates"], "body_bytes": 757, "body_sha256": "sha256:5ec4733347ec5f1825ff77491d3f3698e80386bf8cb28291a855f080e7d8b22f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:coordinates", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/coordinates/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1022303130331221-2300213200322220-3201122010311000-3321110313100221-3131323322013103-1303321103022032-0232221333301332-1033023321332233", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["coordinates"], "schema_version": 1, "sections": [{"aliases": ["coordinates latitude"], "anchor": "schema-coordinates--latitude", "description": "Latitude. Latitude of the site location.", "document_id": "xcsh-docs:data-sources:site:properties:coordinates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["coordinates", "latitude"], "syntax": "attribute", "type": "number"}, {"aliases": ["coordinates longitude"], "anchor": "schema-coordinates--longitude", "description": "Longitude. Longitude of site location.", "document_id": "xcsh-docs:data-sources:site:properties:coordinates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["coordinates", "longitude"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/coordinates/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Coordinates of the site which provides the site physical location.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# coordinates

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- coordinates

<a id="section"></a>

Type: `"single"`. Computed.

Coordinates of the site which provides the site physical location.

## Direct properties

<a id="schema-coordinates--latitude"></a>

### latitude property

Type: `"number"`. Computed.

Latitude. Latitude of the site location.

<a id="schema-coordinates--longitude"></a>

### longitude property

Type: `"number"`. Computed.

Longitude. Longitude of site location.
