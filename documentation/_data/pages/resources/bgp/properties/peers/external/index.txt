---
page_title: "peers.external"
subcategory: ""
description: "External BGP Peer parameters."
xcsh_docs: {"aliases": ["peers external"], "body_bytes": 15865, "body_sha256": "sha256:c2ecc72e288ee992c0f67ff0c4388548eece5bfe2482791c0cc6275d0ecb2c23", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "xcsh-docs:resources:bgp:properties:peers:external:family_inet", "xcsh-docs:resources:bgp:properties:peers:external:from_site", "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "xcsh-docs:resources:bgp:properties:peers:external:interface", "xcsh-docs:resources:bgp:properties:peers:external:interface_list", "xcsh-docs:resources:bgp:properties:peers:external:no_authentication"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "documentation/resources/bgp/properties/peers/external/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [{"anchor": "schema-peers--external--address", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,default_gateway", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,disable_spec", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,external_connector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address_ipv6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,default_gateway_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address_ipv6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,disable_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address_ipv6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,from_site_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address_ipv6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--address_ipv6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--md5_auth_key", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:md5_auth_key,no_authentication", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:external_connector,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:subnet_begin_offset,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_v6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site_v6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_begin_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:subnet_begin_offset_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:external_connector,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:subnet_begin_offset,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "schema-peers--external--subnet_end_offset_v6", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:subnet_begin_offset_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,default_gateway", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,disable_spec", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,external_connector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,default_gateway_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,disable_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,from_site_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,disable_spec", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,disable_spec", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,external_connector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,disable_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,disable_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_v6,from_site_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_v6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,external_connector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,external_connector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,external_connector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:external_connector,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:external_connector,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:external_connector,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_spec,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:external_connector,from_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site,subnet_begin_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site,subnet_end_offset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:address_ipv6,from_site_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:default_gateway_v6,from_site_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:disable_v6,from_site_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site_v6,subnet_begin_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:from_site_v6,subnet_end_offset_v6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:interface,interface_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:interface,interface_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external:ConflictingObjectAttributes:md5_auth_key,no_authentication", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:no_authentication", "type": "conflicts"}, {"anchor": "schema-peers--external--asn", "enforcement": "provider-schema", "group": "peers.external:RequiredObjectAttributes:asn,port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "requires"}, {"anchor": "schema-peers--external--port", "enforcement": "provider-schema", "group": "peers.external:RequiredObjectAttributes:asn,port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external"], "schema_version": 1, "sections": [{"aliases": ["peers external address"], "anchor": "schema-peers--external--address", "description": "Exclusive with Specify IPv4 peer address.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "address"], "syntax": "attribute", "type": "string"}, {"aliases": ["peers external address ipv6"], "anchor": "schema-peers--external--address_ipv6", "description": "Exclusive with Specify peer IPv6 address.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "address_ipv6"], "syntax": "attribute", "type": "string"}, {"aliases": ["peers external asn"], "anchor": "schema-peers--external--asn", "description": "Configure the autonomous system number (ASN) for the external BGP peer.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["peers external default gateway"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "default_gateway"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external default gateway v6"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "default_gateway_v6"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external disable v6"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "disable_v6"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external external connector"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "external_connector"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external family inet"], "anchor": "section", "description": "Parameters for inet family.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "peers.external.family_inet:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external.family_inet:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable", "type": "conflicts"}], "schema_path": ["peers", "external", "family_inet"], "syntax": "block", "type": "object"}, {"aliases": ["peers external from site"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "from_site"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external from site v6"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "from_site_v6"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external interface"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:interface", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-peers--external--interface--name", "enforcement": "provider-schema", "group": "peers.external.interface:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:interface", "type": "requires"}], "schema_path": ["peers", "external", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["peers external interface list"], "anchor": "section", "description": "List of network interfaces.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "peers.external.interface_list:RequiredObjectAttributes:interfaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list:interfaces", "type": "requires"}], "schema_path": ["peers", "external", "interface_list"], "syntax": "block", "type": "object"}, {"aliases": ["peers external md5 auth key"], "anchor": "schema-peers--external--md5_auth_key", "description": "Exclusive with MD5 key for protecting BGP Sessions (RFC 2385)", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "md5_auth_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials", "peers external no authentication"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:no_authentication", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "no_authentication"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external port"], "anchor": "schema-peers--external--port", "description": "Peer TCP port number.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["peers external subnet begin offset"], "anchor": "schema-peers--external--subnet_begin_offset", "description": "Exclusive with Calculate peer address using offset from the beginning of the subnet.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "subnet_begin_offset"], "syntax": "attribute", "type": "number"}, {"aliases": ["peers external subnet begin offset v6"], "anchor": "schema-peers--external--subnet_begin_offset_v6", "description": "Exclusive with Calculate peer address using offset from the beginning of the subnet.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "subnet_begin_offset_v6"], "syntax": "attribute", "type": "number"}, {"aliases": ["peers external subnet end offset"], "anchor": "schema-peers--external--subnet_end_offset", "description": "Exclusive with Calculate peer address using offset from the end of the subnet.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "subnet_end_offset"], "syntax": "attribute", "type": "number"}, {"aliases": ["peers external subnet end offset v6"], "anchor": "schema-peers--external--subnet_end_offset_v6", "description": "Exclusive with Calculate peer address using offset from the end of the subnet.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "subnet_end_offset_v6"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "External BGP Peer parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["bgpCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- peers.external

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

External BGP Peer. External BGP Peer parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn",
    "port"),
  validators.ConflictingObjectAttributes("address",
    "default_gateway"),
  validators.ConflictingObjectAttributes("address",
    "disable_spec"),
  validators.ConflictingObjectAttributes("address",
    "external_connector"),
  validators.ConflictingObjectAttributes("address",
    "from_site"),
  validators.ConflictingObjectAttributes("address",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("address",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "default_gateway_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_gateway",
    "external_connector"),
  validators.ConflictingObjectAttributes("default_gateway",
    "from_site"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("disable_spec",
    "external_connector"),
  validators.ConflictingObjectAttributes("disable_spec",
    "from_site"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("disable_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("external_connector",
    "from_site"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("interface",
    "interface_list"),
  validators.ConflictingObjectAttributes("md5_auth_key",
    "no_authentication"),
  validators.ConflictingObjectAttributes("subnet_begin_offset",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("subnet_begin_offset_v6",
    "subnet_end_offset_v6")}
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
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

Terraform syntax:

```terraform
external {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-peers--external--address"></a>

### address property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-peers--external--address_ipv6"></a>

### address_ipv6 property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-peers--external--asn"></a>

### asn property

Type: `"number"`. Optional.

ASN. Autonomous System Number for BGP peer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/default_gateway/): complete subsection reference.

- [default_gateway_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/default_gateway_v6/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/disable_spec/): complete subsection reference.

- [disable_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/disable_v6/): complete subsection reference.

- [external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/external_connector/): complete subsection reference.

- [family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/): complete subsection reference.

- [from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/from_site/): complete subsection reference.

- [from_site_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/from_site_v6/): complete subsection reference.

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/interface/): complete subsection reference.

- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/interface_list/): complete subsection reference.

<a id="schema-peers--external--md5_auth_key"></a>

### md5_auth_key property

Type: `"string"`. Optional.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/no_authentication/): complete subsection reference.

<a id="schema-peers--external--port"></a>

### port property

Type: `"number"`. Optional.

Peer Port. Peer TCP port number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-peers--external--subnet_begin_offset"></a>

### subnet_begin_offset property

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-peers--external--subnet_begin_offset_v6"></a>

### subnet_begin_offset_v6 property

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-peers--external--subnet_end_offset"></a>

### subnet_end_offset property

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-peers--external--subnet_end_offset_v6"></a>

### subnet_end_offset_v6 property

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```
