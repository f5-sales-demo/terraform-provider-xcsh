---
page_title: "xcsh_external_connector reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector reference."
---

# xcsh_external_connector reference

<a id="canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc6eb83e0cb1bd0e1b09fc420287da095ffc9c85959bdfa81c124d95964d8559"></a>

## Property reference — Property reference / ae3da8957787 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- Property reference

<a id="canonical-30cdfc65d47fc6766bbb914d6ccd9bfce5c842123095afe2ac29d0b6460e6159"></a>

## Direct properties — Property reference / ae3da8957787 / 3

<a id="canonical-be15536c969c16bb54bb0432caf0a8d48a5dc39a6e42e91fdb4703819bfc1139"></a>

<a id="canonical-d38cda88700e4e40d68452a86f41e4879e86c54ced582a70784392f6e4871ea3"></a>

## annotations property — Property reference / ae3da8957787 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [ce_site_reference](resources--external_connector--reference--group-001.md#canonical-b5605bd81ac2bdcfd00cea3b165d31cb83536580ef1c1ca87a76b83a579196ba): complete subsection reference.

<a id="canonical-c04b5c3ff5fcc61833284fe9917b6966b70d4338a0d236558929289f723ca380"></a>

<a id="canonical-22d73705233dded324b0fe7422c89925f7a516372ff448c25e142dd50a51a0c6"></a>

## description property — Property reference / ae3da8957787 / 5

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-49e9e2cc534f24c103de8bfc453ccd28db47aaed0e7c9962ede58c298e09fbab"></a>

<a id="canonical-d3bfd61d38cdf49b41bb47f4f20c0900d9291fa6d69fb59358583026faf233fd"></a>

## disable property — Property reference / ae3da8957787 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd): complete subsection reference.

<a id="canonical-5581ba6be4450f59f066a491e7a76d77542b3c114dbd919cacaef04d01a09a61"></a>

<a id="canonical-cb706d4ed9d30b6504f7bc4f53963dd6ccecc04fa7b4997719d58bc7f1997e28"></a>

## id property — Property reference / ae3da8957787 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f): complete subsection reference.

<a id="canonical-6ef040153b6ab245f47521d9b15e886afa6fadd53a6109b3158838d3e6c8f7d1"></a>

<a id="canonical-46adf5e12a2b92b2032fbaf541925fa5f6d91018bc2f9b2e19cc7f9461b40fb9"></a>

## labels property — Property reference / ae3da8957787 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-a45d3ee742178980e6504d6e673ae643edc7ffe56dbb5fdf0b26028ab63428da"></a>

<a id="canonical-1fff952ab5073c9f112055418916c761c29b3a9dd6086a05229a30c06ba04330"></a>

## name property — Property reference / ae3da8957787 / 9

Type: `"string"`. Required.

Name of the External Connector. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-4fdc876ef145fa7b051af050fc9d06cd05e760684b9d343ec9e5eb7f54561760"></a>

<a id="canonical-4e65dad117de89a1dc29e78179100fe269d42fe4ab3de25635efa4fcde9d32b1"></a>

## namespace property — Property reference / ae3da8957787 / 10

Type: `"string"`. Required.

Namespace where the External Connector is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [timeouts](resources--external_connector--reference--group-001.md#canonical-8bd988d4690c7f02565463b9dc88a47891aad4a4f2df07104bbdced24386fed3): complete subsection reference.

<a id="canonical-57b70aed1843b2109fc844a1aff0d8251fb0c4f2998a096fa995bc0feb5a793d"></a>

## All schema paths — Property reference / ae3da8957787 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--external_connector--reference--group-001.md#canonical-be15536c969c16bb54bb0432caf0a8d48a5dc39a6e42e91fdb4703819bfc1139) |
| `ce_site_reference` | [ce_site_reference](resources--external_connector--reference--group-001.md#canonical-8b817c971f344a578415ae13b5e16424e78bad304d0dff5e30a42db4fd743d89) |
| `ce_site_reference.name` | [ce_site_reference.name](resources--external_connector--reference--group-001.md#canonical-aff7eff974935b094bd5d10de2fc8d76ea861a2221af0399da4aeb168fac2221) |
| `ce_site_reference.namespace` | [ce_site_reference.namespace](resources--external_connector--reference--group-001.md#canonical-a12e5d7a862243a13acf97983af3c2281c1d63e7c6b76bd70503eb795bbb3d99) |
| `ce_site_reference.tenant` | [ce_site_reference.tenant](resources--external_connector--reference--group-001.md#canonical-ef7003037a50e2b1cad455b3deaf868a09a803a1d32f9453a2c52fa9854ddce8) |
| `description` | [description](resources--external_connector--reference--group-001.md#canonical-c04b5c3ff5fcc61833284fe9917b6966b70d4338a0d236558929289f723ca380) |
| `disable` | [disable](resources--external_connector--reference--group-001.md#canonical-49e9e2cc534f24c103de8bfc453ccd28db47aaed0e7c9962ede58c298e09fbab) |
| `gre` | [gre](resources--external_connector--reference--group-001.md#canonical-ec5b41687c6fdd48898e24d484622f0a4022cd3927fe09253dbfa9a25b1030c2) |
| `gre.gre_parameters` | [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-64e6ef0f3c25e75f8f927ea76df07f99cb5d43b7e0f026f8ed7540982442a38d) |
| `gre.gre_parameters.peer_ip_address` | [gre.gre_parameters.peer_ip_address](resources--external_connector--reference--group-001.md#canonical-e4d737384fa8b0bcf5aa14ab654bc0c65764b88a21433c9702faf66c79ebfeee) |
| `gre.gre_parameters.peer_ip_address.addr` | [gre.gre_parameters.peer_ip_address.addr](resources--external_connector--reference--group-001.md#canonical-311d9de5e4bc7e09a021490be33d8557818954db8ff84b491726d9dbe36ace81) |
| `gre.gre_parameters.segment` | [gre.gre_parameters.segment](resources--external_connector--reference--group-001.md#canonical-3e2e95a53b9ab2457745869811f08e62b90b9fd5298bcca6d9a3e9e44cc6ab26) |
| `gre.gre_parameters.segment.refs` | [gre.gre_parameters.segment.refs](resources--external_connector--reference--group-001.md#canonical-04030ecf0a387a5cbaaa9bb55aa5f5a5ae8f4e9c0a89da51671fa804d150da33) |
| `gre.gre_parameters.segment.refs.kind` | [gre.gre_parameters.segment.refs.kind](resources--external_connector--reference--group-001.md#canonical-1333324091d3afd6819ee02c71b0ebc8af9fa496a9eb61dc715b4c7d65b588aa) |
| `gre.gre_parameters.segment.refs.name` | [gre.gre_parameters.segment.refs.name](resources--external_connector--reference--group-001.md#canonical-c4e5f72bc9b9d1a99af77839ac729baa49819ca6adc943a81c2500b729034ddd) |
| `gre.gre_parameters.segment.refs.namespace` | [gre.gre_parameters.segment.refs.namespace](resources--external_connector--reference--group-001.md#canonical-093ca8e0cd2945eb9c39c6d6967f7c0e728e0dda9da59069a6b4a4d4df16621d) |
| `gre.gre_parameters.segment.refs.tenant` | [gre.gre_parameters.segment.refs.tenant](resources--external_connector--reference--group-001.md#canonical-2c90f3234ebf88d070dfb636dfc5cf59104a525bc719b22c90d8dce0d6714b21) |
| `gre.gre_parameters.segment.refs.uid` | [gre.gre_parameters.segment.refs.uid](resources--external_connector--reference--group-001.md#canonical-26b0f03c4e079c018c133fbcfb16dcb37a1587b3ba05c2fae6a8f6c3f184ac00) |
| `gre.gre_parameters.site_local_inside_network` | [gre.gre_parameters.site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-53b3e8db28cce54e1e366cabeac054aa4c4fbbf31876ee1ec40898749f49f8be) |
| `gre.gre_parameters.site_local_network` | [gre.gre_parameters.site_local_network](resources--external_connector--reference--group-001.md#canonical-7e6ae8912f45974fa189f6def040682903d9881cf3489400208164b850d2ed11) |
| `gre.gre_parameters.tunnel_eps` | [gre.gre_parameters.tunnel_eps](resources--external_connector--reference--group-001.md#canonical-514a0a02c3076f2bece21d2b24acea3cac303e4f2a12a516bc5f35262ffe3894) |
| `gre.gre_parameters.tunnel_eps.interface` | [gre.gre_parameters.tunnel_eps.interface](resources--external_connector--reference--group-001.md#canonical-14ee6b2bdd03bb8c60ff2468db911faacabab3dc5470caa6c463d80e3f2bc760) |
| `gre.gre_parameters.tunnel_eps.local_tunnel_ip` | [gre.gre_parameters.tunnel_eps.local_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-ee8b6d1a0e1e5415835e619e14d3ff64924383e8964626176ce33609816de184) |
| `gre.gre_parameters.tunnel_eps.node` | [gre.gre_parameters.tunnel_eps.node](resources--external_connector--reference--group-001.md#canonical-f989db9568c795f56f4c361bf71cb4e71f17bd3a00efc0d3b196270db2b93702) |
| `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` | [gre.gre_parameters.tunnel_eps.remote_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-28bd71373d0a6adc47bae093a711f18ddd6bd39036a343d7c7421399fdccad91) |
| `gre.gre_parameters.tunnel_mtu` | [gre.gre_parameters.tunnel_mtu](resources--external_connector--reference--group-001.md#canonical-3f4f1b3bad586094c0f1bfa58fe374e028a45db1a4ae3a826cc256642dd003f0) |
| `id` | [id](resources--external_connector--reference--group-001.md#canonical-5581ba6be4450f59f066a491e7a76d77542b3c114dbd919cacaef04d01a09a61) |
| `ipsec` | [ipsec](resources--external_connector--reference--group-001.md#canonical-dfc0889217a19b43f0553b214c05cddffc24c4022d573ab47564e9f8c7b3c274) |
| `ipsec.ike_parameters` | [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-3c62ee0ae86dc6c377a980107e7a5a56303bb54444026536ac2cf7b10937a7e1) |
| `ipsec.ike_parameters.dpd_disabled` | [ipsec.ike_parameters.dpd_disabled](resources--external_connector--reference--group-001.md#canonical-00e9b3cc17eecc44019b05f001df886defb77c5cb0024e37f6968ff4735240df) |
| `ipsec.ike_parameters.dpd_keep_alive_timer` | [ipsec.ike_parameters.dpd_keep_alive_timer](resources--external_connector--reference--group-001.md#canonical-50c99f02732e3f1e3659d0e65fb5230581f90b83906ec7cc37f7c5771065c19f) |
| `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` | [ipsec.ike_parameters.dpd_keep_alive_timer.timeout](resources--external_connector--reference--group-001.md#canonical-d4aa570322be0ec08eb33327a72963da4e4947a6c1a20fc191fbf6df1c95c4de) |
| `ipsec.ike_parameters.ike_phase1_profile` | [ipsec.ike_parameters.ike_phase1_profile](resources--external_connector--reference--group-001.md#canonical-f3f0315ca251f809d0865977215e71a329cb718ea8b3effd0a0c499f66a5df94) |
| `ipsec.ike_parameters.ike_phase1_profile.name` | [ipsec.ike_parameters.ike_phase1_profile.name](resources--external_connector--reference--group-001.md#canonical-61a203a6295ce20af732a7fe3062c73a6836f13807e8ecadb7b7e17de4ae7074) |
| `ipsec.ike_parameters.ike_phase1_profile.namespace` | [ipsec.ike_parameters.ike_phase1_profile.namespace](resources--external_connector--reference--group-001.md#canonical-70a630e82ce48f1e11439bbd3130739939a957fd413ed800c940786144325fb4) |
| `ipsec.ike_parameters.ike_phase1_profile.tenant` | [ipsec.ike_parameters.ike_phase1_profile.tenant](resources--external_connector--reference--group-001.md#canonical-1b529b4a73df80a6877126fc37afc23f168e38765e4cd7d14a834f8f73420d43) |
| `ipsec.ike_parameters.ike_phase2_profile` | [ipsec.ike_parameters.ike_phase2_profile](resources--external_connector--reference--group-001.md#canonical-397b0d1f957990dc55b64f378c747267af0fefc2aca4094ec28b69ec6001971c) |
| `ipsec.ike_parameters.ike_phase2_profile.name` | [ipsec.ike_parameters.ike_phase2_profile.name](resources--external_connector--reference--group-001.md#canonical-fb99de8e48139d1ff7690a3e9198f0fedf396432a667c1d07805cf38fcda5561) |
| `ipsec.ike_parameters.ike_phase2_profile.namespace` | [ipsec.ike_parameters.ike_phase2_profile.namespace](resources--external_connector--reference--group-001.md#canonical-3dbd73e7ba00241cd518ae40190f1257a64e9c688c9ae2912da415001a50e750) |
| `ipsec.ike_parameters.ike_phase2_profile.tenant` | [ipsec.ike_parameters.ike_phase2_profile.tenant](resources--external_connector--reference--group-001.md#canonical-510c320b6ba8cc495b0e973b2f6b067841e359e86ba6b311aea1025b67095eb5) |
| `ipsec.ike_parameters.initiator` | [ipsec.ike_parameters.initiator](resources--external_connector--reference--group-001.md#canonical-57475fc460a488546086cdd316e227c77ace23df0cb284d3f2a2ddd6dae64df1) |
| `ipsec.ike_parameters.responder` | [ipsec.ike_parameters.responder](resources--external_connector--reference--group-001.md#canonical-f7c347a562b47bff17d7b2374a2ce8170e90c2b88fa178355e052f5996a742cb) |
| `ipsec.ike_parameters.rm_hostname` | [ipsec.ike_parameters.rm_hostname](resources--external_connector--reference--group-001.md#canonical-96895b7360915c43436899a1bf617d2cdcd1ac6f3ed7e2017a84cf0c2717e36c) |
| `ipsec.ike_parameters.rm_ip_address` | [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-2cdd768d78de91eada626cb38484c306399f8a0a83ecd99d2652b0f1344c1e28) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack` | [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-cdd2756aec91743601fe54687779fb7965de2b6d3334ae98d01cea3d94c2d2cf) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](resources--external_connector--reference--group-001.md#canonical-73f7a69498b308df8d0db9609c425630a16f9319fefc8280404bd7eb4ed4c534) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr](resources--external_connector--reference--group-001.md#canonical-19f4d94efe103c644efd76796f663ce0fada12a401f92386d9acada98e1ed469) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](resources--external_connector--reference--group-001.md#canonical-df6dd0bcb3b9e2ab6920319d975a04c41794c5446a6b53dfdfe8279878d1d73a) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr](resources--external_connector--reference--group-001.md#canonical-bf2a9f83f7eb9171e7f2845fcf84c0ee724ddb034a778d321fc98a6e38c765b4) |
| `ipsec.ike_parameters.rm_ip_address.ipv4` | [ipsec.ike_parameters.rm_ip_address.ipv4](resources--external_connector--reference--group-001.md#canonical-a1150b1c6dced2f7790fff94fb79942a4337f176f79fb316a2bb5dd509409cdf) |
| `ipsec.ike_parameters.rm_ip_address.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.ipv4.addr](resources--external_connector--reference--group-001.md#canonical-ea98b8512ee9494ac06c9e5b8844103d78647be65840bb2f20e6c9dd4abee196) |
| `ipsec.ike_parameters.rm_ip_address.ipv6` | [ipsec.ike_parameters.rm_ip_address.ipv6](resources--external_connector--reference--group-001.md#canonical-91fdefcc357761f773ccaddbc3eb878dc5b1b1d123b98ce2ce6032bf213a3d8a) |
| `ipsec.ike_parameters.rm_ip_address.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.ipv6.addr](resources--external_connector--reference--group-001.md#canonical-7f3dcb6dad9ac5581849aee5b8e93603d5b7ea57939805da4fad6de4be441d23) |
| `ipsec.ike_parameters.use_default_local_ike_id` | [ipsec.ike_parameters.use_default_local_ike_id](resources--external_connector--reference--group-001.md#canonical-c2c716f82791e96277d3ad3a96f53c5141f3dbe7267d03b14ce88a68085f2566) |
| `ipsec.ike_parameters.use_default_remote_ike_id` | [ipsec.ike_parameters.use_default_remote_ike_id](resources--external_connector--reference--group-001.md#canonical-47832d8c329cad275e636058a62b00c43e212d2f8a6a4b0e79f4ee455f9c4b71) |
| `ipsec.ipsec_tunnel_parameters` | [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-8097254c42baa047ba2309465700e13c66f7c050fc7d96627b623a0b1c26b22b) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address` | [ipsec.ipsec_tunnel_parameters.peer_ip_address](resources--external_connector--reference--group-001.md#canonical-6a8906815ece31bb9e8a30ef421a3072ff12339b29f57d93945ad99c4a613a04) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` | [ipsec.ipsec_tunnel_parameters.peer_ip_address.addr](resources--external_connector--reference--group-001.md#canonical-ded6f0ba8e572e9c5fd2f4fbca77f65ca3cd0fdcf8d197e99dfd5be8d3d130eb) |
| `ipsec.ipsec_tunnel_parameters.psk` | [ipsec.ipsec_tunnel_parameters.psk](resources--external_connector--reference--group-001.md#canonical-9cf524d47b3417be89b057571ab1b5cbf0f6966071c67b2d25a7a8bd19bb3783) |
| `ipsec.ipsec_tunnel_parameters.segment` | [ipsec.ipsec_tunnel_parameters.segment](resources--external_connector--reference--group-001.md#canonical-e88f0b0dd5befb5d2b12e38119b98940b630b5ad022d96277929073d386e3393) |
| `ipsec.ipsec_tunnel_parameters.segment.refs` | [ipsec.ipsec_tunnel_parameters.segment.refs](resources--external_connector--reference--group-001.md#canonical-555a6f193404836cbdd8b441ba51cc922673309f7fc315c15f71c400bde7e516) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.kind` | [ipsec.ipsec_tunnel_parameters.segment.refs.kind](resources--external_connector--reference--group-001.md#canonical-42496342f52ff797d51d3ab0ac3f19cb3a682e9ad241c75b2df9448eba848fc0) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.name` | [ipsec.ipsec_tunnel_parameters.segment.refs.name](resources--external_connector--reference--group-001.md#canonical-e489dea88b0a7e46800f1922af1355b9404154a7264b4716b38cee1c453b5ba7) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` | [ipsec.ipsec_tunnel_parameters.segment.refs.namespace](resources--external_connector--reference--group-001.md#canonical-c34b3c8656d87485447ab1e2afc0e5b848416d88ec9f973d34420864f0359578) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` | [ipsec.ipsec_tunnel_parameters.segment.refs.tenant](resources--external_connector--reference--group-001.md#canonical-785166c272cc7f419fde4b41c5e9cd871e28beab07afdf6282925785b6a2778a) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.uid` | [ipsec.ipsec_tunnel_parameters.segment.refs.uid](resources--external_connector--reference--group-001.md#canonical-770dc08d61ebd19be38a5c98e99b090e3729a9c612726b12ae5ca9cc37cd8316) |
| `ipsec.ipsec_tunnel_parameters.site_local_inside_network` | [ipsec.ipsec_tunnel_parameters.site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-8ebe502a3087ab50af93bbdef610ffde25bcce52ec418d5c8ae3cc8375607265) |
| `ipsec.ipsec_tunnel_parameters.site_local_network` | [ipsec.ipsec_tunnel_parameters.site_local_network](resources--external_connector--reference--group-001.md#canonical-417b98afa9243e7b29a906920f68b5849497ce7eb3589a607723d6e8ec8d8ec2) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps` | [ipsec.ipsec_tunnel_parameters.tunnel_eps](resources--external_connector--reference--group-001.md#canonical-9248c99c3892c9921f189bd592865d618cae2f7c7ee93d7410d6229165534457) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.interface](resources--external_connector--reference--group-001.md#canonical-59eeadf8038dc06fd8df7a119750f25fab0943fbae797dc445fc1e5746c56e1e) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-64d58e550549301d4cfff5df8ee8922fe7dcebd142ec13373a1308b52472ec1c) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.node](resources--external_connector--reference--group-001.md#canonical-d0ffdcb7d49f7a2653d01eebd08d52863b640d731fad16a1b266990ddeeb1e43) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-4dbfb576f379651f167d357f99ac111a1b25b56df3f4b90bbe9822335ddc2e9d) |
| `ipsec.ipsec_tunnel_parameters.tunnel_mtu` | [ipsec.ipsec_tunnel_parameters.tunnel_mtu](resources--external_connector--reference--group-001.md#canonical-a2f3d6e0ccd8628bac0aaa475a3b9b8473ae227de5e563c0e6ca3750c9cf5287) |
| `labels` | [labels](resources--external_connector--reference--group-001.md#canonical-6ef040153b6ab245f47521d9b15e886afa6fadd53a6109b3158838d3e6c8f7d1) |
| `name` | [name](resources--external_connector--reference--group-001.md#canonical-a45d3ee742178980e6504d6e673ae643edc7ffe56dbb5fdf0b26028ab63428da) |
| `namespace` | [namespace](resources--external_connector--reference--group-001.md#canonical-4fdc876ef145fa7b051af050fc9d06cd05e760684b9d343ec9e5eb7f54561760) |
| `timeouts` | [timeouts](resources--external_connector--reference--group-001.md#canonical-8ab666f16aed900de56738881281d1346ffbedc2e2bad6ff556aa34959ba8aa9) |
| `timeouts.create` | [timeouts.create](resources--external_connector--reference--group-001.md#canonical-f1df989ba763320d9ae1ae80fc88973e456ba3f07b0c837b5d2f98ed3086b238) |
| `timeouts.delete` | [timeouts.delete](resources--external_connector--reference--group-001.md#canonical-abe14e4807a57173ec311db4db8c2c3222eac1019711a8b8c104c4793d4c4762) |
| `timeouts.read` | [timeouts.read](resources--external_connector--reference--group-001.md#canonical-8666fea6bce4ca8188a95e0a58a8e11ee89e18ffb26e61a6d10bdd5e713ec3a0) |
| `timeouts.update` | [timeouts.update](resources--external_connector--reference--group-001.md#canonical-8412ff3f1101dbd97068ed423dab71718285da0c42638b2af4f2ba4f605cfc96) |

<a id="canonical-e59415f91597f542c4fddb89bd5a5059c70190e5a3446d979f32d2ccc6142ce2"></a>

## Next pages — Property reference / ae3da8957787 / 12

- [ce_site_reference](resources--external_connector--reference--group-001.md#canonical-b5605bd81ac2bdcfd00cea3b165d31cb83536580ef1c1ca87a76b83a579196ba)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [timeouts](resources--external_connector--reference--group-001.md#canonical-8bd988d4690c7f02565463b9dc88a47891aad4a4f2df07104bbdced24386fed3)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-b5605bd81ac2bdcfd00cea3b165d31cb83536580ef1c1ca87a76b83a579196ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddcd104e86ab4ddc9453d218c416bdb96bf4a11dc2263c1c51b25e097eb28092"></a>

## ce_site_reference — ce_site_reference / 18f102cc0006 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- ce_site_reference

<a id="canonical-8b817c971f344a578415ae13b5e16424e78bad304d0dff5e30a42db4fd743d89"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
ce_site_reference {
  # Configure direct properties listed below.
}
```

<a id="canonical-25f6e0b8c3c2545837e74e23ed746cbd5eb4b349b955e08da08020c8036719fc"></a>

## Direct properties — ce_site_reference / 18f102cc0006 / 3

<a id="canonical-aff7eff974935b094bd5d10de2fc8d76ea861a2221af0399da4aeb168fac2221"></a>

<a id="canonical-4639e0377b3b882442fd9e6d2ae48baf35e74544166100309e688bd182842869"></a>

## name property — ce_site_reference / 18f102cc0006 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-a12e5d7a862243a13acf97983af3c2281c1d63e7c6b76bd70503eb795bbb3d99"></a>

<a id="canonical-672af5bda519a3abeb1d16924a82c11b2bf518006410c7b1576481b16e33b31e"></a>

## namespace property — ce_site_reference / 18f102cc0006 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-ef7003037a50e2b1cad455b3deaf868a09a803a1d32f9453a2c52fa9854ddce8"></a>

<a id="canonical-d6907c896e9e8ff91311355cc8e1801f351d113b48863c9b9bdefb18fddc759a"></a>

## tenant property — ce_site_reference / 18f102cc0006 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-c76fe4db05c8976a2f294649154e849dc3e49f76cc5ec30f38c02b68d55870b7"></a>

## Next pages — ce_site_reference / 18f102cc0006 / 7

- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb7741b80dae937ca93e6bd914902935a36538c99e0fff7a743729aea1bb675b"></a>

## gre — gre / 45cbf636b2ad / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- gre

<a id="canonical-ec5b41687c6fdd48898e24d484622f0a4022cd3927fe09253dbfa9a25b1030c2"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: gre, ipsec\] GRE. External Connector with GRE tunnel.

Upstream description:

External Connector with GRE tunnel.

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

OneOf alternatives in this subsection:

- [gre](resources--external_connector--reference--group-001.md#canonical-ec5b41687c6fdd48898e24d484622f0a4022cd3927fe09253dbfa9a25b1030c2)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-dfc0889217a19b43f0553b214c05cddffc24c4022d573ab47564e9f8c7b3c274)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
gre {
  # Configure direct properties listed below.
}
```

<a id="canonical-e4bf17b369bdf70b4d541dbbbc31c024f3dc6179e71e9c0dec1dc36d5ca5f7cd"></a>

## Direct properties — gre / 45cbf636b2ad / 3

- [gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60): complete subsection reference.

<a id="canonical-cb63515412a22e6fd5402458120a569cba7748a7a9a1e0aa196c42772476815b"></a>

## Next pages — gre / 45cbf636b2ad / 4

- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aeef2b13e3c64aad8d0c69d3fd7fb7b7ca13ede91b4c2455a6fa94498fd4fb5c"></a>

## gre.gre_parameters — gre.gre_parameters / 05b469e2fb9f / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- gre.gre_parameters

<a id="canonical-64e6ef0f3c25e75f8f927ea76df07f99cb5d43b7e0f026f8ed7540982442a38d"></a>

Type: `"object"`. single nested block, Optional.

GRE configuration parameters required for GRE Connection type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tunnel_eps",
    "tunnel_mtu"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_inside_network"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-tunnel_source_vn": "[\"segment\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
gre_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-9f41942c51a60cef08268cbd71c205a2c5ea7a5d69c1398cbf6eb21f4c52a598"></a>

## Direct properties — gre.gre_parameters / 05b469e2fb9f / 3

- [peer_ip_address](resources--external_connector--reference--group-001.md#canonical-b5a895d475718bdfb43c1822a5023d44d97ef847195db3deb1dd21f239dbc65a): complete subsection reference.

- [segment](resources--external_connector--reference--group-001.md#canonical-cf1534113d37737a79aca028992e94d64d6d27d734f5a0a4ae03d1a860ecc32c): complete subsection reference.

- [site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-3bce261607e25d801f473f4971a4edf03f257c14bd64081dfeeca379918d4dc5): complete subsection reference.

- [site_local_network](resources--external_connector--reference--group-001.md#canonical-dac7baacd4768f6be0669dae9420bfb0aa19371b4148eaa287f510d6a0a0e986): complete subsection reference.

- [tunnel_eps](resources--external_connector--reference--group-001.md#canonical-c458d97d312583a4266d128cf2550f62a8c47f8a78130301031e1db347539066): complete subsection reference.

<a id="canonical-3f4f1b3bad586094c0f1bfa58fe374e028a45db1a4ae3a826cc256642dd003f0"></a>

<a id="canonical-64a1a507ffc4f36d870463c0240149cb0e1dbeba8464c3d29b5ec6518ca320c2"></a>

## tunnel_mtu property — gre.gre_parameters / 05b469e2fb9f / 4

Type: `"number"`. Optional.

Configure MTU for the GRE tunnel interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(512, 1370),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1370,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 512
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```

<a id="canonical-ec55c80579dd85cc3517ba96599499c20ffa608c23722a19742537838d02bea2"></a>

## Next pages — gre.gre_parameters / 05b469e2fb9f / 5

- [gre.gre_parameters.peer_ip_address](resources--external_connector--reference--group-001.md#canonical-b5a895d475718bdfb43c1822a5023d44d97ef847195db3deb1dd21f239dbc65a)
- [gre.gre_parameters.segment](resources--external_connector--reference--group-001.md#canonical-cf1534113d37737a79aca028992e94d64d6d27d734f5a0a4ae03d1a860ecc32c)
- [gre.gre_parameters.site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-3bce261607e25d801f473f4971a4edf03f257c14bd64081dfeeca379918d4dc5)
- [gre.gre_parameters.site_local_network](resources--external_connector--reference--group-001.md#canonical-dac7baacd4768f6be0669dae9420bfb0aa19371b4148eaa287f510d6a0a0e986)
- [gre.gre_parameters.tunnel_eps](resources--external_connector--reference--group-001.md#canonical-c458d97d312583a4266d128cf2550f62a8c47f8a78130301031e1db347539066)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-b5a895d475718bdfb43c1822a5023d44d97ef847195db3deb1dd21f239dbc65a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d0fe5df7b95ddb567a14bdef5ccf7fd9a9e6fbfeea75b1ad854f5336ba8f2af"></a>

## gre.gre_parameters.peer_ip_address — gre.gre_parameters.peer_ip_address / 8adc8274e61f / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- gre.gre_parameters.peer_ip_address

<a id="canonical-e4d737384fa8b0bcf5aa14ab654bc0c65764b88a21433c9702faf66c79ebfeee"></a>

Type: `"object"`. single nested block, Optional.

IPv4 Address. IPv4 Address in dot-decimal notation.

Upstream description:

IPv4 Address in dot-decimal notation.

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
peer_ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-86f1dec7fe24ca42235748c02ab7a4d532c3ce0fdbaa7407edf7a71d8545a69d"></a>

## Direct properties — gre.gre_parameters.peer_ip_address / 8adc8274e61f / 3

<a id="canonical-311d9de5e4bc7e09a021490be33d8557818954db8ff84b491726d9dbe36ace81"></a>

<a id="canonical-e6aed6c83a3e1cb62ce8fd70066ce29cdfbb4a87a714240dd11052432f441a54"></a>

## addr property — gre.gre_parameters.peer_ip_address / 8adc8274e61f / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-6fc2774149235f7f774916b84a779d47d49822c9bd979a6ff8290e03be1cccd9"></a>

## Next pages — gre.gre_parameters.peer_ip_address / 8adc8274e61f / 5

- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-cf1534113d37737a79aca028992e94d64d6d27d734f5a0a4ae03d1a860ecc32c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bbcd868ab931ec1bbb537d2f4a43e4f87b916d73b5e728706e2505a2f70f903"></a>

## gre.gre_parameters.segment — gre.gre_parameters.segment / f495b1e4d5ac / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- gre.gre_parameters.segment

<a id="canonical-3e2e95a53b9ab2457745869811f08e62b90b9fd5298bcca6d9a3e9e44cc6ab26"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-42abe125b9d57f89f98191f17231f8c5a631b36e07dce97ebe8eab46cbdeeb23"></a>

## Direct properties — gre.gre_parameters.segment / f495b1e4d5ac / 3

- [refs](resources--external_connector--reference--group-001.md#canonical-1c2456321bfbb3f011d47f07b9298304ef284fe920467dfb23980019e3ce0425): complete subsection reference.

<a id="canonical-4f848714f9bca419cc6ac6c2f3ac33756f96f3caab9b3a401f64dea6d67f85a5"></a>

## Next pages — gre.gre_parameters.segment / f495b1e4d5ac / 4

- [gre.gre_parameters.segment.refs](resources--external_connector--reference--group-001.md#canonical-1c2456321bfbb3f011d47f07b9298304ef284fe920467dfb23980019e3ce0425)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-1c2456321bfbb3f011d47f07b9298304ef284fe920467dfb23980019e3ce0425"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5681d18d6e89eecbf13b89b6c8b26eea4d57eccbc190c0b8cfdbbef38de2add2"></a>

## gre.gre_parameters.segment.refs — gre.gre_parameters.segment.refs / 9d639916663e / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- [gre.gre_parameters.segment](resources--external_connector--reference--group-001.md#canonical-cf1534113d37737a79aca028992e94d64d6d27d734f5a0a4ae03d1a860ecc32c)
- gre.gre_parameters.segment.refs

<a id="canonical-04030ecf0a387a5cbaaa9bb55aa5f5a5ae8f4e9c0a89da51671fa804d150da33"></a>

Type: `"object"`. list nested block, Optional.

Segment. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-681b527d54ddf84023c0b616bf74349ee1d118463fd6d8e8226c84588224f4d3"></a>

## Direct properties — gre.gre_parameters.segment.refs / 9d639916663e / 3

<a id="canonical-1333324091d3afd6819ee02c71b0ebc8af9fa496a9eb61dc715b4c7d65b588aa"></a>

<a id="canonical-51d926e929de717b126efe6584cae79ae75dfb3996e95ad3014db91c8476ed69"></a>

## kind property — gre.gre_parameters.segment.refs / 9d639916663e / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-c4e5f72bc9b9d1a99af77839ac729baa49819ca6adc943a81c2500b729034ddd"></a>

<a id="canonical-5ac9608379dab527eb18c762d31f8281b5ed0b2305303285ed57cda04139aec8"></a>

## name property — gre.gre_parameters.segment.refs / 9d639916663e / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-093ca8e0cd2945eb9c39c6d6967f7c0e728e0dda9da59069a6b4a4d4df16621d"></a>

<a id="canonical-78ff10f9042fee730e5c525f79d89987af72772167e26bb31cac0db3e014d55d"></a>

## namespace property — gre.gre_parameters.segment.refs / 9d639916663e / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="canonical-2c90f3234ebf88d070dfb636dfc5cf59104a525bc719b22c90d8dce0d6714b21"></a>

<a id="canonical-a6afff2058a4d1d08750cb3fb7a54324b1ec29064570fb4c165b2372cc60cdbd"></a>

## tenant property — gre.gre_parameters.segment.refs / 9d639916663e / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-26b0f03c4e079c018c133fbcfb16dcb37a1587b3ba05c2fae6a8f6c3f184ac00"></a>

<a id="canonical-682f2c9320e3a834347c4e7786ce126443bae21990228d11ebab36e4daf16319"></a>

## uid property — gre.gre_parameters.segment.refs / 9d639916663e / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-7b8f67baa70aea4a384244777c42c365bd2dabc961e062648ff71f12b121e107"></a>

## Next pages — gre.gre_parameters.segment.refs / 9d639916663e / 9

- [gre.gre_parameters.segment](resources--external_connector--reference--group-001.md#canonical-cf1534113d37737a79aca028992e94d64d6d27d734f5a0a4ae03d1a860ecc32c)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-3bce261607e25d801f473f4971a4edf03f257c14bd64081dfeeca379918d4dc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-618077dd5e60473f700b6a164a945481969adb1ba11090b10ab328341e9ae46a"></a>

## gre.gre_parameters.site_local_inside_network — gre.gre_parameters.site_local_inside_network / a9b607a08aa8 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- gre.gre_parameters.site_local_inside_network

<a id="canonical-53b3e8db28cce54e1e366cabeac054aa4c4fbbf31876ee1ec40898749f49f8be"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
site_local_inside_network = {}
```

<a id="canonical-78e765e071bd68b5919ce5daccfa483d318ddf1cf3ce6887bdff65f07b356ba5"></a>

## Direct properties — gre.gre_parameters.site_local_inside_network / a9b607a08aa8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-caff245a72e14b3540a3d08a9f5fc8910ef6b65d37c77e63791ee32ee3985acb"></a>

## Next pages — gre.gre_parameters.site_local_inside_network / a9b607a08aa8 / 4

- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-dac7baacd4768f6be0669dae9420bfb0aa19371b4148eaa287f510d6a0a0e986"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f4e48f3579174b443dc56e2d6494991505a2c9aaa478ad1b69c023c07b000ed"></a>

## gre.gre_parameters.site_local_network — gre.gre_parameters.site_local_network / e78198410df1 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- gre.gre_parameters.site_local_network

<a id="canonical-7e6ae8912f45974fa189f6def040682903d9881cf3489400208164b850d2ed11"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
site_local_network = {}
```

<a id="canonical-989c89fea1ea1d15e99b3bb4541bc05eb2cc2bc66fbfde53a982bdaf9eb77ad1"></a>

## Direct properties — gre.gre_parameters.site_local_network / e78198410df1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e235d0daf8a30d31e82fe97dba3a8e7a249591610278d6b3f28bfc9c08a271d"></a>

## Next pages — gre.gre_parameters.site_local_network / e78198410df1 / 4

- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-c458d97d312583a4266d128cf2550f62a8c47f8a78130301031e1db347539066"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a60f8132adfc14533e9a7347a9fa22e5e6937155c82423c5dd8038c88623f367"></a>

## gre.gre_parameters.tunnel_eps — gre.gre_parameters.tunnel_eps / 1bf1912d457c / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [gre](resources--external_connector--reference--group-001.md#canonical-4f0205f939f13dbaece45757bac67d3396f58aca26296ab2d6a4adbf1565f3dd)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- gre.gre_parameters.tunnel_eps

<a id="canonical-514a0a02c3076f2bece21d2b24acea3cac303e4f2a12a516bc5f35262ffe3894"></a>

Type: `"object"`. list nested block, Optional.

Configure tunnel parameters, source, destination, IP addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("interface",
    "local_tunnel_ip",
    "node",
    "remote_tunnel_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tunnel_eps {
  # Configure direct properties listed below.
}
```

<a id="canonical-10a656115fcf42af1f82c20877738bd2313e8f34e738251d16a95a083bc08a81"></a>

## Direct properties — gre.gre_parameters.tunnel_eps / 1bf1912d457c / 3

<a id="canonical-14ee6b2bdd03bb8c60ff2468db911faacabab3dc5470caa6c463d80e3f2bc760"></a>

<a id="canonical-7edbc5993f055d69347b53a4a735ce745c44ddaf0e4eec2d60d9029b438980b8"></a>

## interface property — gre.gre_parameters.tunnel_eps / 1bf1912d457c / 4

Type: `"string"`. Optional.

For the chosen node, specify the interface that will be the tunnel source.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-ee8b6d1a0e1e5415835e619e14d3ff64924383e8964626176ce33609816de184"></a>

<a id="canonical-c2960c73f0f7a2efaa4fac0848976bc10bae809fd9fc1a212ba8d4e53db27487"></a>

## local_tunnel_ip property — gre.gre_parameters.tunnel_eps / 1bf1912d457c / 5

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the local tunnel IP Address i.e. The IP address of the
tunnel on the CE node itself and a subnet prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-f989db9568c795f56f4c361bf71cb4e71f17bd3a00efc0d3b196270db2b93702"></a>

<a id="canonical-5376bf8eee143c60f8387061a9e9892a3a90a0405e6895602c5b7cc64500d658"></a>

## node property — gre.gre_parameters.tunnel_eps / 1bf1912d457c / 6

Type: `"string"`. Optional.

CE site is composed of multiple nodes. Choose a node that will be part of this external connection.

Upstream description:

A CE site is composed of multiple nodes. Choose a node that will be part of this external
connection.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-28bd71373d0a6adc47bae093a711f18ddd6bd39036a343d7c7421399fdccad91"></a>

<a id="canonical-578b95f5f8db4f6a7fc49780f6cbbe85606907f037f5429f04d314e4abaa0860"></a>

## remote_tunnel_ip property — gre.gre_parameters.tunnel_eps / 1bf1912d457c / 7

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the remote tunnel IP Address i.e. The IP address of the
tunnel on the remote gateway and a subnet prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-d0bbf6a78211059a1ea5e0f01fa971eda293a770327509e82cb93c6d9dfb64f5"></a>

## Next pages — gre.gre_parameters.tunnel_eps / 1bf1912d457c / 8

- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-d1c51f9994d82007ad05590d153a57e463aa2f7882c9a9e1506d2ad8f7748c60)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d6a0f3ec06074388fafc854438c85032c10bab1e17fdf293e694719cc54ca16"></a>

## ipsec — ipsec / e99bcc707427 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- ipsec

<a id="canonical-dfc0889217a19b43f0553b214c05cddffc24c4022d573ab47564e9f8c7b3c274"></a>

Type: `"object"`. single nested block, Optional.

IPsec. External Connector with IPsec tunnel.

Upstream description:

External Connector with IPsec tunnel.

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
ipsec {
  # Configure direct properties listed below.
}
```

<a id="canonical-d22670c31d81e4761942b4bccf7d725ad6569dba7ba62a2268a799cbdf2f6fd4"></a>

## Direct properties — ipsec / e99bcc707427 / 3

- [ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947): complete subsection reference.

- [ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18): complete subsection reference.

<a id="canonical-f63007d8f32e3f860e9e0da74ba6b50cc7f7642f5ebff6606a5542c930418e03"></a>

## Next pages — ipsec / e99bcc707427 / 4

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f2dc617d5594540350a28c74a128949674874d0e21e912ba778e9b0c44ac5b9"></a>

## ipsec.ike_parameters — ipsec.ike_parameters / f9c8f8f1ad09 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- ipsec.ike_parameters

<a id="canonical-3c62ee0ae86dc6c377a980107e7a5a56303bb54444026536ac2cf7b10937a7e1"></a>

Type: `"object"`. single nested block, Optional.

IKE configuration parameters required for IPsec Connection type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dpd_disabled",
    "dpd_keep_alive_timer"),
  validators.ConflictingObjectAttributes("initiator",
    "responder"),
  validators.ConflictingObjectAttributes("rm_hostname",
    "rm_ip_address"),
  validators.ConflictingObjectAttributes("rm_hostname",
    "use_default_remote_ike_id"),
  validators.ConflictingObjectAttributes("rm_ip_address",
    "use_default_remote_ike_id")}
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
  "x-ves-oneof-field-dpd_choice": "[\"dpd_disabled\",\"dpd_keep_alive_timer\"]",
  "x-ves-oneof-field-local_ike_id": "[\"use_default_local_ike_id\"]",
  "x-ves-oneof-field-mode_choice": "[\"initiator\",\"responder\"]",
  "x-ves-oneof-field-remote_ike_id": "[\"rm_hostname\",\"rm_ip_address\",\"use_default_remote_ike_id\"]"
}
```

Terraform syntax:

```terraform
ike_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4cdb2dd00f773ea8664ccbb2a8a0374bd4c21c2dde69bf46fe30a0393ad453c"></a>

## Direct properties — ipsec.ike_parameters / f9c8f8f1ad09 / 3

- [dpd_disabled](resources--external_connector--reference--group-001.md#canonical-55e50928152db659a6dd3ccdc9bf5fc1090caf82f74f5cb6dfd75fe4ae1833fb): complete subsection reference.

- [dpd_keep_alive_timer](resources--external_connector--reference--group-001.md#canonical-2183b47ac4357fbf244db6db4680110fc9d6aaca673b75d1dcecb10d28caee84): complete subsection reference.

- [ike_phase1_profile](resources--external_connector--reference--group-001.md#canonical-25b986b17d792e868c5529eb854be71555f407c3736e6634a1419eeb53eeea07): complete subsection reference.

- [ike_phase2_profile](resources--external_connector--reference--group-001.md#canonical-d5d7be9c13fd8b0f36422c1a9ca871c15815ebf84c295790160dec912e760a2b): complete subsection reference.

- [initiator](resources--external_connector--reference--group-001.md#canonical-32661be4a27e01408921dcebd3f404ad4f630846e8f2115256a4e183aadc4453): complete subsection reference.

- [responder](resources--external_connector--reference--group-001.md#canonical-0e8c76f03f21680414dbb746e97020cfb5acdfb47d658198487e69ef1367874f): complete subsection reference.

<a id="canonical-96895b7360915c43436899a1bf617d2cdcd1ac6f3ed7e2017a84cf0c2717e36c"></a>

<a id="canonical-879633309f565d4f6aa289bcaecff1d56ac24d028c3c6ee3cee541ca120a2b8c"></a>

## rm_hostname property — ipsec.ike_parameters / f9c8f8f1ad09 / 4

Type: `"string"`. Optional.

Exclusive with \[rm\_ip\_address use\_default\_remote\_ike\_id\] Configure an hostname Remote IKE
ID.

Upstream description:

Exclusive with \[rm\_ip\_address use\_default\_remote\_ike\_id\] Configure an hostname Remote IKE
ID.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877): complete subsection reference.

- [use_default_local_ike_id](resources--external_connector--reference--group-001.md#canonical-a3fa656d6fba43edf283715378e0d7aac81479a58524dc3f62e076d06d4d7716): complete subsection reference.

- [use_default_remote_ike_id](resources--external_connector--reference--group-001.md#canonical-bea6ba76bc63b979fc55b350dc019447a1c35d5abe8d24fe8eb0b742276844ea): complete subsection reference.

<a id="canonical-ef11e825258c1076ab30b5a4f4007c4c3c48459230a2db808d47cf71d9c0a69e"></a>

## Next pages — ipsec.ike_parameters / f9c8f8f1ad09 / 5

- [ipsec.ike_parameters.dpd_disabled](resources--external_connector--reference--group-001.md#canonical-55e50928152db659a6dd3ccdc9bf5fc1090caf82f74f5cb6dfd75fe4ae1833fb)
- [ipsec.ike_parameters.dpd_keep_alive_timer](resources--external_connector--reference--group-001.md#canonical-2183b47ac4357fbf244db6db4680110fc9d6aaca673b75d1dcecb10d28caee84)
- [ipsec.ike_parameters.ike_phase1_profile](resources--external_connector--reference--group-001.md#canonical-25b986b17d792e868c5529eb854be71555f407c3736e6634a1419eeb53eeea07)
- [ipsec.ike_parameters.ike_phase2_profile](resources--external_connector--reference--group-001.md#canonical-d5d7be9c13fd8b0f36422c1a9ca871c15815ebf84c295790160dec912e760a2b)
- [ipsec.ike_parameters.initiator](resources--external_connector--reference--group-001.md#canonical-32661be4a27e01408921dcebd3f404ad4f630846e8f2115256a4e183aadc4453)
- [ipsec.ike_parameters.responder](resources--external_connector--reference--group-001.md#canonical-0e8c76f03f21680414dbb746e97020cfb5acdfb47d658198487e69ef1367874f)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- [ipsec.ike_parameters.use_default_local_ike_id](resources--external_connector--reference--group-001.md#canonical-a3fa656d6fba43edf283715378e0d7aac81479a58524dc3f62e076d06d4d7716)
- [ipsec.ike_parameters.use_default_remote_ike_id](resources--external_connector--reference--group-001.md#canonical-bea6ba76bc63b979fc55b350dc019447a1c35d5abe8d24fe8eb0b742276844ea)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-55e50928152db659a6dd3ccdc9bf5fc1090caf82f74f5cb6dfd75fe4ae1833fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-454d45c1f0dad0ab16eece178c2b57ab910dbf970d9c20b11e54761da17d6443"></a>

## ipsec.ike_parameters.dpd_disabled — ipsec.ike_parameters.dpd_disabled / 37766c7b87fc / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.dpd_disabled

<a id="canonical-00e9b3cc17eecc44019b05f001df886defb77c5cb0024e37f6968ff4735240df"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
dpd_disabled = {}
```

<a id="canonical-20239e8e12a212109c83384faa27ccb28831332d1c3674f0feff8a22afaf4b61"></a>

## Direct properties — ipsec.ike_parameters.dpd_disabled / 37766c7b87fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-563f06f15c5828ca38ee6ebd960ca9e28683bc3c9429472f5083e8ac50c3433b"></a>

## Next pages — ipsec.ike_parameters.dpd_disabled / 37766c7b87fc / 4

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-2183b47ac4357fbf244db6db4680110fc9d6aaca673b75d1dcecb10d28caee84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-649a9c2bee7353241c804ee22ba30d360400e9d0e805d68c3fa0de2f29ad58ca"></a>

## ipsec.ike_parameters.dpd_keep_alive_timer — ipsec.ike_parameters.dpd_keep_alive_timer / feff6be185f0 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.dpd_keep_alive_timer

<a id="canonical-50c99f02732e3f1e3659d0e65fb5230581f90b83906ec7cc37f7c5771065c19f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dpd keep alive timer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("timeout")}
```

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
dpd_keep_alive_timer {
  # Configure direct properties listed below.
}
```

<a id="canonical-ec51ee784368a042caaba99571a02ecfceecb8fb844f8dff4cc336150c76fd95"></a>

## Direct properties — ipsec.ike_parameters.dpd_keep_alive_timer / feff6be185f0 / 3

<a id="canonical-d4aa570322be0ec08eb33327a72963da4e4947a6c1a20fc191fbf6df1c95c4de"></a>

<a id="canonical-71756b796e97e7120a87a4d987d39b7815d422b47120c1f00433806360049db0"></a>

## timeout property — ipsec.ike_parameters.dpd_keep_alive_timer / feff6be185f0 / 4

Type: `"number"`. Optional.

Keepalive Timer. Operation timeout duration

Upstream description:

Operation timeout duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-28f16014ff2e6256a142de342a7cc4f877e2f3c176dcede6caa8702591ccd2ba"></a>

## Next pages — ipsec.ike_parameters.dpd_keep_alive_timer / feff6be185f0 / 5

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-25b986b17d792e868c5529eb854be71555f407c3736e6634a1419eeb53eeea07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47ac7354f4d8568388588568324b39e8ee18e9252394d9ceeaf3cec9e0f3fcdf"></a>

## ipsec.ike_parameters.ike_phase1_profile — ipsec.ike_parameters.ike_phase1_profile / bcc3d4379bfe / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.ike_phase1_profile

<a id="canonical-f3f0315ca251f809d0865977215e71a329cb718ea8b3effd0a0c499f66a5df94"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
ike_phase1_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0df68e42064c380713f30a90f6b5c2e62cd79d7d9304dc42285b098cf03c7765"></a>

## Direct properties — ipsec.ike_parameters.ike_phase1_profile / bcc3d4379bfe / 3

<a id="canonical-61a203a6295ce20af732a7fe3062c73a6836f13807e8ecadb7b7e17de4ae7074"></a>

<a id="canonical-48b6183d078c5760a8a884612e0b109dc3b064e50ca4117f61debb4cdd8f22cb"></a>

## name property — ipsec.ike_parameters.ike_phase1_profile / bcc3d4379bfe / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-70a630e82ce48f1e11439bbd3130739939a957fd413ed800c940786144325fb4"></a>

<a id="canonical-111e32f683c1cbadc68b6ef0448a71fbf5f02daecf06f51aab99eb12eaea4153"></a>

## namespace property — ipsec.ike_parameters.ike_phase1_profile / bcc3d4379bfe / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1b529b4a73df80a6877126fc37afc23f168e38765e4cd7d14a834f8f73420d43"></a>

<a id="canonical-cae7ab42660ea9a8769573b46d39afdf066a2cdca0dcf5660022a52df45737ad"></a>

## tenant property — ipsec.ike_parameters.ike_phase1_profile / bcc3d4379bfe / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-9b83c3aed1ea15a539ad953d4076d0aeadec6e2785428a712a341ea579ad0ea5"></a>

## Next pages — ipsec.ike_parameters.ike_phase1_profile / bcc3d4379bfe / 7

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-d5d7be9c13fd8b0f36422c1a9ca871c15815ebf84c295790160dec912e760a2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fc22914362d14956cd1fa84633eda09028c2ba179d409a6403c830a1fe318ae"></a>

## ipsec.ike_parameters.ike_phase2_profile — ipsec.ike_parameters.ike_phase2_profile / 5582a3c87d67 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.ike_phase2_profile

<a id="canonical-397b0d1f957990dc55b64f378c747267af0fefc2aca4094ec28b69ec6001971c"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
ike_phase2_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-7efd97327cefb58b0d64c9941df8d6fb77bf5215e4253100c912fc2ca5d70a27"></a>

## Direct properties — ipsec.ike_parameters.ike_phase2_profile / 5582a3c87d67 / 3

<a id="canonical-fb99de8e48139d1ff7690a3e9198f0fedf396432a667c1d07805cf38fcda5561"></a>

<a id="canonical-adbda399fbbd251e6ff1aa52920d7de9c1bb8aa30a40c4c99d5ed66d005741f1"></a>

## name property — ipsec.ike_parameters.ike_phase2_profile / 5582a3c87d67 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3dbd73e7ba00241cd518ae40190f1257a64e9c688c9ae2912da415001a50e750"></a>

<a id="canonical-a54fcdc60096f78b985402d069c15ee932fe5857e8c20a6daa8cde9cd9ff64e5"></a>

## namespace property — ipsec.ike_parameters.ike_phase2_profile / 5582a3c87d67 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-510c320b6ba8cc495b0e973b2f6b067841e359e86ba6b311aea1025b67095eb5"></a>

<a id="canonical-0db31bb169d8b902e15583142ac59417b2e37e54bf92c9cbe1071d045eb0c236"></a>

## tenant property — ipsec.ike_parameters.ike_phase2_profile / 5582a3c87d67 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-73ff6b35693bb3dc766439a66d47e2286ac8a7551186a2f38cd841666d50c5cb"></a>

## Next pages — ipsec.ike_parameters.ike_phase2_profile / 5582a3c87d67 / 7

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-32661be4a27e01408921dcebd3f404ad4f630846e8f2115256a4e183aadc4453"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71896c7a68123f70d064b04ee530b8020c97a35b3ffab29a69cf82d4ce8310e1"></a>

## ipsec.ike_parameters.initiator — ipsec.ike_parameters.initiator / dfb52fdabdec / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.initiator

<a id="canonical-57475fc460a488546086cdd316e227c77ace23df0cb284d3f2a2ddd6dae64df1"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
initiator = {}
```

<a id="canonical-e7b8634cc3d619f971382aa9eb8d84a0438bbc33fd73ecb93662d06b5333e4d6"></a>

## Direct properties — ipsec.ike_parameters.initiator / dfb52fdabdec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffa3ee64c8393290a698930b6f1c7d443edcade46d016920a289c81ddd905c6e"></a>

## Next pages — ipsec.ike_parameters.initiator / dfb52fdabdec / 4

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-0e8c76f03f21680414dbb746e97020cfb5acdfb47d658198487e69ef1367874f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efc46c3e3910fc4f176f0673e2892d2ddbc13e0a6c5a84745b52879f6a076a3d"></a>

## ipsec.ike_parameters.responder — ipsec.ike_parameters.responder / feccbebc9d68 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.responder

<a id="canonical-f7c347a562b47bff17d7b2374a2ce8170e90c2b88fa178355e052f5996a742cb"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
responder = {}
```

<a id="canonical-263887a73225d6041f80fd87dd06ae08010559d4dc993619b203dd1ab581f5f5"></a>

## Direct properties — ipsec.ike_parameters.responder / feccbebc9d68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7924a7faf0dd9a592142768d001590231bf44809d1071c75fd31bf5ee0a8a42f"></a>

## Next pages — ipsec.ike_parameters.responder / feccbebc9d68 / 4

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75561528f441b6c0b67d092613f70c11de65cd40dd917b7216ad3a326c994a23"></a>

## ipsec.ike_parameters.rm_ip_address — ipsec.ike_parameters.rm_ip_address / d1a027462f13 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.rm_ip_address

<a id="canonical-2cdd768d78de91eada626cb38484c306399f8a0a83ecd99d2652b0f1344c1e28"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
rm_ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-a341362cb3e78122e7d531a5c79b3109d1ed2e3f870b7b2957bd34320c958118"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address / d1a027462f13 / 3

- [dual_stack](resources--external_connector--reference--group-001.md#canonical-f097090df0e1581225f3e1f8eb50b7c8d9ba800da87bdd0307daa84366cd657b): complete subsection reference.

- [ipv4](resources--external_connector--reference--group-001.md#canonical-9094b52166b21a64a6fe23a08113a85b65de50523ae9fed693dcdf7c436738e9): complete subsection reference.

- [ipv6](resources--external_connector--reference--group-001.md#canonical-1775b471ae165404735cc63e57894f20c40956438faa640d307d3ed6918d542c): complete subsection reference.

<a id="canonical-ba4c6874118644d2eb360692d9cbceff105e0d7d1b4ff7144a42aa2a92908c49"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address / d1a027462f13 / 4

- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-f097090df0e1581225f3e1f8eb50b7c8d9ba800da87bdd0307daa84366cd657b)
- [ipsec.ike_parameters.rm_ip_address.ipv4](resources--external_connector--reference--group-001.md#canonical-9094b52166b21a64a6fe23a08113a85b65de50523ae9fed693dcdf7c436738e9)
- [ipsec.ike_parameters.rm_ip_address.ipv6](resources--external_connector--reference--group-001.md#canonical-1775b471ae165404735cc63e57894f20c40956438faa640d307d3ed6918d542c)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-f097090df0e1581225f3e1f8eb50b7c8d9ba800da87bdd0307daa84366cd657b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bc405c1c8a29e7d61297e1124f0d7c506a7607a699152679cfabb00ef705974"></a>

## ipsec.ike_parameters.rm_ip_address.dual_stack — ipsec.ike_parameters.rm_ip_address.dual_stack / 0b947b24e11c / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- ipsec.ike_parameters.rm_ip_address.dual_stack

<a id="canonical-cdd2756aec91743601fe54687779fb7965de2b6d3334ae98d01cea3d94c2d2cf"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-798e121030e53a846739b54e11dae1ebabc2548d9dd9e9dd88b4fd5a86d4db87"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.dual_stack / 0b947b24e11c / 3

- [ipv4](resources--external_connector--reference--group-001.md#canonical-44f3626f9b8bfb568a84748793f5daa6184b5288f7122c7f0a128b14fc1f71d4): complete subsection reference.

- [ipv6](resources--external_connector--reference--group-001.md#canonical-2f469dc0c5656b78f708730418aa4819e6913504b22bdc9e3c7486f2ec7f4023): complete subsection reference.

<a id="canonical-1920932324da849fdfcda70de528fdbe31bf5ba7f7bb43160d0d3628babea7a6"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.dual_stack / 0b947b24e11c / 4

- [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](resources--external_connector--reference--group-001.md#canonical-44f3626f9b8bfb568a84748793f5daa6184b5288f7122c7f0a128b14fc1f71d4)
- [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](resources--external_connector--reference--group-001.md#canonical-2f469dc0c5656b78f708730418aa4819e6913504b22bdc9e3c7486f2ec7f4023)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-44f3626f9b8bfb568a84748793f5daa6184b5288f7122c7f0a128b14fc1f71d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecbad0b74238e834fb8e71c73f29db07c98e0ffc20554ca77203d39c2a32b14f"></a>

## ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 8015dc26834a / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-f097090df0e1581225f3e1f8eb50b7c8d9ba800da87bdd0307daa84366cd657b)
- ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4

<a id="canonical-73f7a69498b308df8d0db9609c425630a16f9319fefc8280404bd7eb4ed4c534"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-eded067e535fea06d35111907e4d9ac6b68b4d7e4cf3f7e93a88150b8b99da3e"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 8015dc26834a / 3

<a id="canonical-19f4d94efe103c644efd76796f663ce0fada12a401f92386d9acada98e1ed469"></a>

<a id="canonical-bb7b4c99d70ab5a3848285ac3b8873369244ead5fea6f3d85a98634d51a5976d"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 8015dc26834a / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-d4c066c0cd82d0c3ffdc4ee1172530cd5303315e3de12316a4669d5f16cde335"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 8015dc26834a / 5

- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-f097090df0e1581225f3e1f8eb50b7c8d9ba800da87bdd0307daa84366cd657b)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-2f469dc0c5656b78f708730418aa4819e6913504b22bdc9e3c7486f2ec7f4023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82348e12252b3fed338079348164bc79d6246692cb5d419dbc1b81d9b88dc935"></a>

## ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 393a49fb9984 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-f097090df0e1581225f3e1f8eb50b7c8d9ba800da87bdd0307daa84366cd657b)
- ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6

<a id="canonical-df6dd0bcb3b9e2ab6920319d975a04c41794c5446a6b53dfdfe8279878d1d73a"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-221c1682e7bbe9c212524a6d906d9f0e2d7333cfcdc0925a9c67e954bf9dc8a0"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 393a49fb9984 / 3

<a id="canonical-bf2a9f83f7eb9171e7f2845fcf84c0ee724ddb034a778d321fc98a6e38c765b4"></a>

<a id="canonical-ef0266b69772da6b50b01b075a432dc11c62def1676d2a6778d820364098ca9b"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 393a49fb9984 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-c8aa11c2e381a38db87d5df3e1767cf1c894baaae4cf7402003518258bd34a4d"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 393a49fb9984 / 5

- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-f097090df0e1581225f3e1f8eb50b7c8d9ba800da87bdd0307daa84366cd657b)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-9094b52166b21a64a6fe23a08113a85b65de50523ae9fed693dcdf7c436738e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfcfe60da955313370eeef7d1eb927c0b7ff41bb20097582a5af57083fca5ffb"></a>

## ipsec.ike_parameters.rm_ip_address.ipv4 — ipsec.ike_parameters.rm_ip_address.ipv4 / 957dcad5fd43 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- ipsec.ike_parameters.rm_ip_address.ipv4

<a id="canonical-a1150b1c6dced2f7790fff94fb79942a4337f176f79fb316a2bb5dd509409cdf"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3cb51bfdba1905e7a6a9611f7cdef38d8e5fc0b0353430d68f9561198c72b776"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.ipv4 / 957dcad5fd43 / 3

<a id="canonical-ea98b8512ee9494ac06c9e5b8844103d78647be65840bb2f20e6c9dd4abee196"></a>

<a id="canonical-24c07b98ab33a343e1e5ebd72832d1ea2943f8344aece58e06113ed14cfb7899"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.ipv4 / 957dcad5fd43 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-31fcd7afd032298635d69fea6be1f488e2f97f3a5c74d984886e87a9a3188a34"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.ipv4 / 957dcad5fd43 / 5

- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-1775b471ae165404735cc63e57894f20c40956438faa640d307d3ed6918d542c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73f3d9ea55d115ed3c075fe6f8f0ce4c7bd33b5bef67123200eae324608dcf2d"></a>

## ipsec.ike_parameters.rm_ip_address.ipv6 — ipsec.ike_parameters.rm_ip_address.ipv6 / f63bd1042e9d / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- ipsec.ike_parameters.rm_ip_address.ipv6

<a id="canonical-91fdefcc357761f773ccaddbc3eb878dc5b1b1d123b98ce2ce6032bf213a3d8a"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3a0295572306d84f7aa6e17cced71575b5ae460c70ebb698ad3209318db01f57"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.ipv6 / f63bd1042e9d / 3

<a id="canonical-7f3dcb6dad9ac5581849aee5b8e93603d5b7ea57939805da4fad6de4be441d23"></a>

<a id="canonical-fdcff62def685afdd1e78a0dd265a135c47ec48e8a2a43a34489bc7c926cd59c"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.ipv6 / f63bd1042e9d / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-90c0c99acab9f2436ea8c2d8d87d2e049ceb8b4c270cbbd91e1974e43cda3852"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.ipv6 / f63bd1042e9d / 5

- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-49e24b2f30739a913e5e10f69bb90a06e10f40a4fa6afbe2d5c10d10255da877)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-a3fa656d6fba43edf283715378e0d7aac81479a58524dc3f62e076d06d4d7716"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-895ab74a43ff9817c21b0a9834154d1196612a281a27aaa9eb0f98bcbcc024bb"></a>

## ipsec.ike_parameters.use_default_local_ike_id — ipsec.ike_parameters.use_default_local_ike_id / 95ae07ca7edb / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.use_default_local_ike_id

<a id="canonical-c2c716f82791e96277d3ad3a96f53c5141f3dbe7267d03b14ce88a68085f2566"></a>

Type: `"object"`. single nested block, Optional.

Enable this option

Upstream description:

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
use_default_local_ike_id {}
```

<a id="canonical-770084720e17da2cf2fb09cca9eaf523b4df32c9a214327355ad97a5a67324f2"></a>

## Direct properties — ipsec.ike_parameters.use_default_local_ike_id / 95ae07ca7edb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d40c228da41db7d9e8fb229e55395307e177227a35e4d75d0f3534d03191487e"></a>

## Next pages — ipsec.ike_parameters.use_default_local_ike_id / 95ae07ca7edb / 4

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-bea6ba76bc63b979fc55b350dc019447a1c35d5abe8d24fe8eb0b742276844ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f64b65403d292afe5f19d2937df0aa2d21608d65e645e158b25c7bdf81ea6c5"></a>

## ipsec.ike_parameters.use_default_remote_ike_id — ipsec.ike_parameters.use_default_remote_ike_id / 25be33714cbe / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- ipsec.ike_parameters.use_default_remote_ike_id

<a id="canonical-47832d8c329cad275e636058a62b00c43e212d2f8a6a4b0e79f4ee455f9c4b71"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
use_default_remote_ike_id = {}
```

<a id="canonical-733bb1968737d7ad3d85db3b5b6e702ce002e9587f0daf335e172ce242c44f02"></a>

## Direct properties — ipsec.ike_parameters.use_default_remote_ike_id / 25be33714cbe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1495ba58c08f8c10bee11338dba9d563fd7c2da3c54decbe85585cacfd5cc973"></a>

## Next pages — ipsec.ike_parameters.use_default_remote_ike_id / 25be33714cbe / 4

- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-9458acac93aa25844673a327d0f7fb4adef28c738602de9b097d51d9ea8cb947)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d1e46e26643b7878daa9fbb3ccd0c89fc8403a3ffd9980a4a35a8faefa8316a"></a>

## ipsec.ipsec_tunnel_parameters — ipsec.ipsec_tunnel_parameters / eb1152a80b10 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- ipsec.ipsec_tunnel_parameters

<a id="canonical-8097254c42baa047ba2309465700e13c66f7c050fc7d96627b623a0b1c26b22b"></a>

Type: `"object"`. single nested block, Optional.

In this section, we will configure the tunnel parameters, source, destination, IP addresses, and
segment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("psk",
    "tunnel_eps",
    "tunnel_mtu"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_inside_network"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-tunnel_source_vn": "[\"segment\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
ipsec_tunnel_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-840f326bbf33bec8809adc4c9bfba00fa3a2ac3acc2858765a1a71d073991eb5"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters / eb1152a80b10 / 3

- [peer_ip_address](resources--external_connector--reference--group-001.md#canonical-c9c5781e64231a9a1e146faf3343caa7d5fc4e6832d318713e5cadff22c223e6): complete subsection reference.

<a id="canonical-9cf524d47b3417be89b057571ab1b5cbf0f6966071c67b2d25a7a8bd19bb3783"></a>

<a id="canonical-e4392a486e29b1786abac896440650eff439c208c1d6855040ebb8b50ea472f5"></a>

## psk property — ipsec.ipsec_tunnel_parameters / eb1152a80b10 / 4

Type: `"string"`. Optional.

The IKE pre-shared key (PSK) is required to ensure the IKE peers can authenticate one another within
IKE phase 1 negotiation.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [segment](resources--external_connector--reference--group-001.md#canonical-7a81dc09a0b81db19bb80874ddca1b09000c94cd6a3497185fb0d5fbf9cda25e): complete subsection reference.

- [site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-be58a97272d554803b733802373c9919fb16eda4f8c877b4cbf04edc46e1c81c): complete subsection reference.

- [site_local_network](resources--external_connector--reference--group-001.md#canonical-267c2c98a5f570b1c7bf6e0fcb45257648526dfff09844bd38ba4c9e1f2ffec6): complete subsection reference.

- [tunnel_eps](resources--external_connector--reference--group-001.md#canonical-5d8a0caefb73859d0ebfea613f4926f992de3020297ecb09ad945cc36ffa29b2): complete subsection reference.

<a id="canonical-a2f3d6e0ccd8628bac0aaa475a3b9b8473ae227de5e563c0e6ca3750c9cf5287"></a>

<a id="canonical-8802fe2b17ec1625f282f1eb852e6233036fa77b3efd10fc8b69591c5624d4db"></a>

## tunnel_mtu property — ipsec.ipsec_tunnel_parameters / eb1152a80b10 / 5

Type: `"number"`. Optional.

The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without
needing to be fragmented.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(512, 1370),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1370,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 512
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```

<a id="canonical-1fd3744bc82c562f674fbbc5c6505d5c63326d6fc066735198afe540e93a3013"></a>

## Next pages — ipsec.ipsec_tunnel_parameters / eb1152a80b10 / 6

- [ipsec.ipsec_tunnel_parameters.peer_ip_address](resources--external_connector--reference--group-001.md#canonical-c9c5781e64231a9a1e146faf3343caa7d5fc4e6832d318713e5cadff22c223e6)
- [ipsec.ipsec_tunnel_parameters.segment](resources--external_connector--reference--group-001.md#canonical-7a81dc09a0b81db19bb80874ddca1b09000c94cd6a3497185fb0d5fbf9cda25e)
- [ipsec.ipsec_tunnel_parameters.site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-be58a97272d554803b733802373c9919fb16eda4f8c877b4cbf04edc46e1c81c)
- [ipsec.ipsec_tunnel_parameters.site_local_network](resources--external_connector--reference--group-001.md#canonical-267c2c98a5f570b1c7bf6e0fcb45257648526dfff09844bd38ba4c9e1f2ffec6)
- [ipsec.ipsec_tunnel_parameters.tunnel_eps](resources--external_connector--reference--group-001.md#canonical-5d8a0caefb73859d0ebfea613f4926f992de3020297ecb09ad945cc36ffa29b2)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-c9c5781e64231a9a1e146faf3343caa7d5fc4e6832d318713e5cadff22c223e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5247d8f8965d8a189efc5565c21c421c7077d449ede6acb7069c63bae9f4c36f"></a>

## ipsec.ipsec_tunnel_parameters.peer_ip_address — ipsec.ipsec_tunnel_parameters.peer_ip_address / e94ea2eeef02 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- ipsec.ipsec_tunnel_parameters.peer_ip_address

<a id="canonical-6a8906815ece31bb9e8a30ef421a3072ff12339b29f57d93945ad99c4a613a04"></a>

Type: `"object"`. single nested block, Optional.

IPv4 Address. IPv4 Address in dot-decimal notation.

Upstream description:

IPv4 Address in dot-decimal notation.

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
peer_ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-c80a0e2d2b314c9174525dbaae1e32f28b830a4cd66b4a81f67c5b9a74eb0e00"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.peer_ip_address / e94ea2eeef02 / 3

<a id="canonical-ded6f0ba8e572e9c5fd2f4fbca77f65ca3cd0fdcf8d197e99dfd5be8d3d130eb"></a>

<a id="canonical-c06017625fc8f94d2fbede7c4aad8cb5758bdb31da08d661e25484aa9bd4c3e7"></a>

## addr property — ipsec.ipsec_tunnel_parameters.peer_ip_address / e94ea2eeef02 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1e7665dec8ce6b86140a39bca66bd0118cf90227eaf474aa1fa007b39712d57d"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.peer_ip_address / e94ea2eeef02 / 5

- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-7a81dc09a0b81db19bb80874ddca1b09000c94cd6a3497185fb0d5fbf9cda25e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09f361f7a870e03e3b4600f89425eab038d758c718122b1733ab9f5e2853d8c8"></a>

## ipsec.ipsec_tunnel_parameters.segment — ipsec.ipsec_tunnel_parameters.segment / 5adffaa7909e / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- ipsec.ipsec_tunnel_parameters.segment

<a id="canonical-e88f0b0dd5befb5d2b12e38119b98940b630b5ad022d96277929073d386e3393"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d65ad2f975e9a5517565fe5ccbc39af7ee98e20d0374ed0a2626d01f7d6d80d"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.segment / 5adffaa7909e / 3

- [refs](resources--external_connector--reference--group-001.md#canonical-433b79da2ca04b18bd54ce45b30d3507f019c714dd252fa05cdbecda9a0952b9): complete subsection reference.

<a id="canonical-56dfbf6b1f65903334f2d81b331f5d0224390ea0b954c41d3236f8ceb2ac2c12"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.segment / 5adffaa7909e / 4

- [ipsec.ipsec_tunnel_parameters.segment.refs](resources--external_connector--reference--group-001.md#canonical-433b79da2ca04b18bd54ce45b30d3507f019c714dd252fa05cdbecda9a0952b9)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-433b79da2ca04b18bd54ce45b30d3507f019c714dd252fa05cdbecda9a0952b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd71c4ccb0f21e5abe270d9c140d7d9af25fb66296df11a1bfc7a664b743d826"></a>

## ipsec.ipsec_tunnel_parameters.segment.refs — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- [ipsec.ipsec_tunnel_parameters.segment](resources--external_connector--reference--group-001.md#canonical-7a81dc09a0b81db19bb80874ddca1b09000c94cd6a3497185fb0d5fbf9cda25e)
- ipsec.ipsec_tunnel_parameters.segment.refs

<a id="canonical-555a6f193404836cbdd8b441ba51cc922673309f7fc315c15f71c400bde7e516"></a>

Type: `"object"`. list nested block, Optional.

Segment. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-4363370fe218b0cb6e1efae5b7132fb08ff4fb4904d06cf87df98566333faa4e"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 3

<a id="canonical-42496342f52ff797d51d3ab0ac3f19cb3a682e9ad241c75b2df9448eba848fc0"></a>

<a id="canonical-5c184ea7b4cda6a0de0801a846c9b7e20b4b71942427e6807474f83591f64348"></a>

## kind property — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-e489dea88b0a7e46800f1922af1355b9404154a7264b4716b38cee1c453b5ba7"></a>

<a id="canonical-a7706b21aef8dae8f6635de2684442277fa04538d22d3d0fc9b93f817f977a01"></a>

## name property — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-c34b3c8656d87485447ab1e2afc0e5b848416d88ec9f973d34420864f0359578"></a>

<a id="canonical-a2d919513febe8b1db2ed852a42a213f5f4eb3e1e0e6916accbdfd2278253077"></a>

## namespace property — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="canonical-785166c272cc7f419fde4b41c5e9cd871e28beab07afdf6282925785b6a2778a"></a>

<a id="canonical-0057931fcfdb60fe277eee1df3e816f56a86db6e87e10124dfe760b14f8866f0"></a>

## tenant property — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-770dc08d61ebd19be38a5c98e99b090e3729a9c612726b12ae5ca9cc37cd8316"></a>

<a id="canonical-7f8c632feb43d3a968a401a424cd1fd1a7bf607f94003c2e48eaeddd3d48f40e"></a>

## uid property — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-9bee97c8d846a219474475f3ccc5741b01017336de2e050a3829fd8a1bad2d4b"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.segment.refs / 1ccbd46d9a9d / 9

- [ipsec.ipsec_tunnel_parameters.segment](resources--external_connector--reference--group-001.md#canonical-7a81dc09a0b81db19bb80874ddca1b09000c94cd6a3497185fb0d5fbf9cda25e)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-be58a97272d554803b733802373c9919fb16eda4f8c877b4cbf04edc46e1c81c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dfd57680eb8ff0ec86cd3aca3de8c1dbaa1b6b2e0205349c412d225a49938c6"></a>

## ipsec.ipsec_tunnel_parameters.site_local_inside_network — ipsec.ipsec_tunnel_parameters.site_local_inside_network / b08bfdaa27b0 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- ipsec.ipsec_tunnel_parameters.site_local_inside_network

<a id="canonical-8ebe502a3087ab50af93bbdef610ffde25bcce52ec418d5c8ae3cc8375607265"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
site_local_inside_network = {}
```

<a id="canonical-58fbcfbe3b917cc5f353abc66bbfe30b619751ce77db308d80d12423f4ef6eaa"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.site_local_inside_network / b08bfdaa27b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-305eb6132dc2d4361d91523e1687d2a3a415eb07e5927089d3bcd68852786141"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.site_local_inside_network / b08bfdaa27b0 / 4

- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-267c2c98a5f570b1c7bf6e0fcb45257648526dfff09844bd38ba4c9e1f2ffec6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d3f877663b9397464bdb4d54096eb570806f085c55537d6d2053208cf03fc5b"></a>

## ipsec.ipsec_tunnel_parameters.site_local_network — ipsec.ipsec_tunnel_parameters.site_local_network / 2b23d31ddfe7 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- ipsec.ipsec_tunnel_parameters.site_local_network

<a id="canonical-417b98afa9243e7b29a906920f68b5849497ce7eb3589a607723d6e8ec8d8ec2"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
site_local_network = {}
```

<a id="canonical-76f47c388b591ec34dfa88cec5de43b8df495345012c87d0dd691697d8388308"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.site_local_network / 2b23d31ddfe7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86f58e378e273d365507749d0461ce89f7672f5ae8f23053374423d64ffca569"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.site_local_network / 2b23d31ddfe7 / 4

- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-5d8a0caefb73859d0ebfea613f4926f992de3020297ecb09ad945cc36ffa29b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c498173428cc2a9b406c69acb77316f31d0ebe72dc3db47d8c9ebac7027db62e"></a>

## ipsec.ipsec_tunnel_parameters.tunnel_eps — ipsec.ipsec_tunnel_parameters.tunnel_eps / 1432eb5de620 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-aa93d639f48fc7527e57b482709cd0edb6907ffbe92a722bf6fc1d592ae8875f)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- ipsec.ipsec_tunnel_parameters.tunnel_eps

<a id="canonical-9248c99c3892c9921f189bd592865d618cae2f7c7ee93d7410d6229165534457"></a>

Type: `"object"`. list nested block, Optional.

Configure tunnel parameters, local and remote IP addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("interface",
    "local_tunnel_ip",
    "node",
    "remote_tunnel_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tunnel_eps {
  # Configure direct properties listed below.
}
```

<a id="canonical-fdc3eb6a5232fccb9556c583a4afca23fbb372bbc924b17f738c64e0f91aa322"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.tunnel_eps / 1432eb5de620 / 3

<a id="canonical-59eeadf8038dc06fd8df7a119750f25fab0943fbae797dc445fc1e5746c56e1e"></a>

<a id="canonical-8f244574aebfe1e4a8de52a377009bfbb691904f2bc1816fa8cbe06ded1d52f2"></a>

## interface property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 1432eb5de620 / 4

Type: `"string"`. Optional.

For the chosen node, specify the interface that will be the tunnel source.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-64d58e550549301d4cfff5df8ee8922fe7dcebd142ec13373a1308b52472ec1c"></a>

<a id="canonical-492b7b5a005939588fc453b9b925d0b2a3b9c17ebd485db29a21b7dd81bb3f4d"></a>

## local_tunnel_ip property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 1432eb5de620 / 5

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the local tunnel IP Address i.e. The IP address of the
tunnel on the CE node itself and a subnet prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-d0ffdcb7d49f7a2653d01eebd08d52863b640d731fad16a1b266990ddeeb1e43"></a>

<a id="canonical-9319ac8e905d5aa7647955edd577dd6fb670fa68979c243d937c192d073a4bb9"></a>

## node property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 1432eb5de620 / 6

Type: `"string"`. Optional.

CE site is composed of multiple nodes. Choose a node that will be part of this external connection.

Upstream description:

A CE site is composed of multiple nodes. Choose a node that will be part of this external
connection.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-4dbfb576f379651f167d357f99ac111a1b25b56df3f4b90bbe9822335ddc2e9d"></a>

<a id="canonical-6abfc0bc2980228896137db695949c9cc9f01b9760eb7312da2030333f906f1c"></a>

## remote_tunnel_ip property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 1432eb5de620 / 7

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the remote tunnel IP Address i.e. The IP address of the
tunnel on the remote gateway and a subnet prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-d580f07ee8fad782dc731eac3687961dd0dc676793854d7b69fb6a59364e400a"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.tunnel_eps / 1432eb5de620 / 8

- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-55da43d47b04220731cd2173fb2ab8d0dcfcad7ccf012b7a7c07e5293bab0c18)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-8bd988d4690c7f02565463b9dc88a47891aad4a4f2df07104bbdced24386fed3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cebdd0116ca73f02b052490d2c5c75cd90cc00b2625f7e75a3ecff94be841fe1"></a>

## timeouts — timeouts / d6295a6d9154 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- timeouts

<a id="canonical-8ab666f16aed900de56738881281d1346ffbedc2e2bad6ff556aa34959ba8aa9"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3cef5ee304b80482b4ae51337faf8095cae52071ec4cc47d99a240528da8bd8f"></a>

## Direct properties — timeouts / d6295a6d9154 / 3

<a id="canonical-f1df989ba763320d9ae1ae80fc88973e456ba3f07b0c837b5d2f98ed3086b238"></a>

<a id="canonical-5bb83b4ae3f7baf7e4538473af979b31805381716e3a918e9e7ff20261474a56"></a>

## create property — timeouts / d6295a6d9154 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-abe14e4807a57173ec311db4db8c2c3222eac1019711a8b8c104c4793d4c4762"></a>

<a id="canonical-83a5fad42fa70286f591c0f91e1bf1a967027cf3940787e60a73f8f3743b8db9"></a>

## delete property — timeouts / d6295a6d9154 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-8666fea6bce4ca8188a95e0a58a8e11ee89e18ffb26e61a6d10bdd5e713ec3a0"></a>

<a id="canonical-73b6b5bdd3bad623b611ce295107444d195204f0b3092d3189c0feff2fd827f3"></a>

## read property — timeouts / d6295a6d9154 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-8412ff3f1101dbd97068ed423dab71718285da0c42638b2af4f2ba4f605cfc96"></a>

<a id="canonical-1753f1d6275cf9bd41df72a09df9b1a75a53821b8f780ead8a8b627784f04176"></a>

## update property — timeouts / d6295a6d9154 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1b74d34df773a0a80dcffe9d98051ac83f95f2d31584e2b687c362eab81a1473"></a>

## Next pages — timeouts / d6295a6d9154 / 8

- [Property reference](resources--external_connector--reference--group-001.md#canonical-e7d57047a482351e27fecdbabb074fcae77652462ab30d4d5f3bcf74ce8b21ee)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
