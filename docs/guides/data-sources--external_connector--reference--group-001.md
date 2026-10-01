---
page_title: "xcsh_external_connector reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector reference."
---

# xcsh_external_connector reference

<a id="canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b63cd1a2cea84b053324cbf621e68b9f83f9a96afa3f224ebdd1994f60694e0"></a>

## Property reference — Property reference / 9ca2569bd122 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- Property reference

<a id="canonical-ade57a74c20a542f5ce8a4930f715ad13c6eed66f63d61042852e4dba80c7ce4"></a>

## Direct properties — Property reference / 9ca2569bd122 / 3

<a id="canonical-a94a1c3dab5fb86b3f95d76f04a6cc966d9f70139757b5b3202d26fa86c7849a"></a>

<a id="canonical-9e012f31ff62d7c27fa563e873833443b1290e0c9b928594c3641b61f0c3b78f"></a>

## annotations property — Property reference / 9ca2569bd122 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [ce_site_reference](data-sources--external_connector--reference--group-001.md#canonical-09dea781a39324e59d78d3c5b6229f1e4279f242425a36e80e87b6d8373ec61c): complete subsection reference.

<a id="canonical-0e0d74b8b39f15ace12bf5a7bfa637049181bf613d0c82830b3b1030cb6a0ecb"></a>

<a id="canonical-d910309e2baedb4baa93da5ca121d0c7deaee98e66999e51a36ee2bb7ca1a5de"></a>

## description property — Property reference / 9ca2569bd122 / 5

Type: `"string"`. Computed.

Description of the ExternalConnector.

Upstream description:

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

- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949): complete subsection reference.

<a id="canonical-36781f916d460866f0bccfa91f791859d4033a915551621e7c6ce3ee735ce302"></a>

<a id="canonical-98bfca62b1e3a32d59b00fc6df2ebdd1383ca4e2c1d6f7b3dbcf3016d1e3ae39"></a>

## id property — Property reference / 9ca2569bd122 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef): complete subsection reference.

<a id="canonical-5ca2721972ae70a29563e0be8e77530d4b1c4c49be18e41235f101b941bf15e1"></a>

<a id="canonical-81100810de36cb7b023f09824d1d6c3122160e43456bc96446cedbd958d6c277"></a>

## labels property — Property reference / 9ca2569bd122 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-12d478a5288698d9a4257f8bfdbe28e535dfb25506fd57220ca7860a03e8bd65"></a>

<a id="canonical-e1b2a4faf153dcaa17c3a5e04dd372507fc1c0accf287f12b909b5a500804f4b"></a>

## name property — Property reference / 9ca2569bd122 / 8

Type: `"string"`. Required.

Name of the ExternalConnector.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-c1353d1649d4f90e20f3acb61737378ab6161d78efccf24984aef715ff6d97d4"></a>

<a id="canonical-c60bad7222a1890b60a6d2572dac71d966b5d69063241c1433e0fc5b0e60b4eb"></a>

## namespace property — Property reference / 9ca2569bd122 / 9

Type: `"string"`. Required.

Namespace where the ExternalConnector exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-2a1496fbfeaf3cf71b8fc34f2727b0232cd2a3862432209c507d11fcdcea8f34"></a>

## All schema paths — Property reference / 9ca2569bd122 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--external_connector--reference--group-001.md#canonical-a94a1c3dab5fb86b3f95d76f04a6cc966d9f70139757b5b3202d26fa86c7849a) |
| `ce_site_reference` | [ce_site_reference](data-sources--external_connector--reference--group-001.md#canonical-912cefb2cae19844958aa9334db2fcf8d965d1736b82eaeee2e40404a4719a5c) |
| `ce_site_reference.name` | [ce_site_reference.name](data-sources--external_connector--reference--group-001.md#canonical-e5eec23648ac135b9f796f42e8c7a8eee0b51c62e22cc913dcb61bd48ab10725) |
| `ce_site_reference.namespace` | [ce_site_reference.namespace](data-sources--external_connector--reference--group-001.md#canonical-da2df306b1419edea322844b645186e3d73f18c74eb0ced1ddd89e7b0e7e1253) |
| `ce_site_reference.tenant` | [ce_site_reference.tenant](data-sources--external_connector--reference--group-001.md#canonical-c24bbdf9580d2a05b6fe74c0de9c128bf8df45aa43bff3ef6df8c2d1d9152b06) |
| `description` | [description](data-sources--external_connector--reference--group-001.md#canonical-0e0d74b8b39f15ace12bf5a7bfa637049181bf613d0c82830b3b1030cb6a0ecb) |
| `gre` | [gre](data-sources--external_connector--reference--group-001.md#canonical-ff62476e01a590e20a11ed2b62b29f1d03c7483558dbd04d2c550e8b8ca9dc22) |
| `gre.gre_parameters` | [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-e6575c113a22db1ca43504b05794881600e9998fb4f96d44e30f29cc6f650594) |
| `gre.gre_parameters.peer_ip_address` | [gre.gre_parameters.peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-55a976bc510e3113f94c59cf30decffeaaa7ce8ca235de5662252c06e7b59ab6) |
| `gre.gre_parameters.peer_ip_address.addr` | [gre.gre_parameters.peer_ip_address.addr](data-sources--external_connector--reference--group-001.md#canonical-705727c5659cfdf62078d15f0dfa590660879f89bdd9f591b6e848f4b22316bd) |
| `gre.gre_parameters.segment` | [gre.gre_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-dfb5d9b018095a4876d25bfed7a9eb2bcd2b4befccc45df5507c0d8f39fa4f67) |
| `gre.gre_parameters.segment.refs` | [gre.gre_parameters.segment.refs](data-sources--external_connector--reference--group-001.md#canonical-3f7c4d99475b45884c182ab6165a424107d867a40e3a4e7c879eca4898f82d10) |
| `gre.gre_parameters.segment.refs.kind` | [gre.gre_parameters.segment.refs.kind](data-sources--external_connector--reference--group-001.md#canonical-4d1c094170ad29a43cf23b4acb7e0fce0a7c182f936f6e45c6d32f04a49e2917) |
| `gre.gre_parameters.segment.refs.name` | [gre.gre_parameters.segment.refs.name](data-sources--external_connector--reference--group-001.md#canonical-a1c4da369e713730fcaabda604e153c4cac406174cf3896a962378ac99a2f836) |
| `gre.gre_parameters.segment.refs.namespace` | [gre.gre_parameters.segment.refs.namespace](data-sources--external_connector--reference--group-001.md#canonical-f03459834916bc458a99759444a4a0e50dfe9e67383f9870bb0bb27c38eb4e72) |
| `gre.gre_parameters.segment.refs.tenant` | [gre.gre_parameters.segment.refs.tenant](data-sources--external_connector--reference--group-001.md#canonical-ee404db6709fa31d9a36a025686b2b0c3dd419f92756790589ed98f915dc9f6e) |
| `gre.gre_parameters.segment.refs.uid` | [gre.gre_parameters.segment.refs.uid](data-sources--external_connector--reference--group-001.md#canonical-740c4b39507e4d01d4da49dba6bbd87e3058fd52aaae8e2ca9e311a82b923ffe) |
| `gre.gre_parameters.site_local_inside_network` | [gre.gre_parameters.site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-38dea2993dfa5ea83b02fb50f2212464c01ace60d8877b83670e7acafb0f2572) |
| `gre.gre_parameters.site_local_network` | [gre.gre_parameters.site_local_network](data-sources--external_connector--reference--group-001.md#canonical-2699bfb6f7577f7e1cb6166c490ca01026936e089ce3ff5c3941906a45eff276) |
| `gre.gre_parameters.tunnel_eps` | [gre.gre_parameters.tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-28fd97bd92d93e588d69f3fabc410050207ede0f2e1807cb29f3d40af13e2b31) |
| `gre.gre_parameters.tunnel_eps.interface` | [gre.gre_parameters.tunnel_eps.interface](data-sources--external_connector--reference--group-001.md#canonical-a65d91b3b23d800c45103b48c47c49c076db7ce85a1eb112306e7695eefda43f) |
| `gre.gre_parameters.tunnel_eps.local_tunnel_ip` | [gre.gre_parameters.tunnel_eps.local_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-1c8db96255256826719b3ed32cab586fc48497adbdc15464cb8c70df9390200c) |
| `gre.gre_parameters.tunnel_eps.node` | [gre.gre_parameters.tunnel_eps.node](data-sources--external_connector--reference--group-001.md#canonical-b4f29d8d16618ba7f1e708ba4b89baf848a5d62b3e348b86bfbdcee7d0bf1bf6) |
| `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` | [gre.gre_parameters.tunnel_eps.remote_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-02b1f127d72056e0cdebe847579e21f3e7803b6c8e1f93be43a3065e620d932f) |
| `gre.gre_parameters.tunnel_mtu` | [gre.gre_parameters.tunnel_mtu](data-sources--external_connector--reference--group-001.md#canonical-1233fdbb7a68df57b95f20e45234cb59eacc93e4db3ad945e4818a533a1a0513) |
| `id` | [id](data-sources--external_connector--reference--group-001.md#canonical-36781f916d460866f0bccfa91f791859d4033a915551621e7c6ce3ee735ce302) |
| `ipsec` | [ipsec](data-sources--external_connector--reference--group-001.md#canonical-d09b756cc40d88ccfd47ab8d95dc4b012620da92abb968542c3335011a0e7861) |
| `ipsec.ike_parameters` | [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-8d6fa631584e216a8a3f68044ebc83a95e67455f512b0448adaa3f2d760e40f2) |
| `ipsec.ike_parameters.dpd_disabled` | [ipsec.ike_parameters.dpd_disabled](data-sources--external_connector--reference--group-001.md#canonical-13958225e44f15c2744f7708aad30f329c867acdf6a6f594f5df74925ee7ccd4) |
| `ipsec.ike_parameters.dpd_keep_alive_timer` | [ipsec.ike_parameters.dpd_keep_alive_timer](data-sources--external_connector--reference--group-001.md#canonical-3e26b634a9c76d9266128338532af832c486f043fcd8e9eca221b3a3d9309d87) |
| `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` | [ipsec.ike_parameters.dpd_keep_alive_timer.timeout](data-sources--external_connector--reference--group-001.md#canonical-1ab0e6975e186637d91682861180a50f5b3d1ed57eaf3297cfe0668df47b2ecc) |
| `ipsec.ike_parameters.ike_phase1_profile` | [ipsec.ike_parameters.ike_phase1_profile](data-sources--external_connector--reference--group-001.md#canonical-1ed4c8d0f51d9b258c3810520833d4e3a9682bf923a70191e33431da7fb5b423) |
| `ipsec.ike_parameters.ike_phase1_profile.name` | [ipsec.ike_parameters.ike_phase1_profile.name](data-sources--external_connector--reference--group-001.md#canonical-03ef8cd99768d1875b4dace0973fd65d74bd14883c8ee75bc7c27b845b34d67e) |
| `ipsec.ike_parameters.ike_phase1_profile.namespace` | [ipsec.ike_parameters.ike_phase1_profile.namespace](data-sources--external_connector--reference--group-001.md#canonical-46c4f25bf74b24e6fe20102ce1315fb32847e94ff6142f2307368bade2b35972) |
| `ipsec.ike_parameters.ike_phase1_profile.tenant` | [ipsec.ike_parameters.ike_phase1_profile.tenant](data-sources--external_connector--reference--group-001.md#canonical-f8a5b5e021a71c60ace9725113ece4f333e993353b6aa4760f78eb6110ad87a2) |
| `ipsec.ike_parameters.ike_phase2_profile` | [ipsec.ike_parameters.ike_phase2_profile](data-sources--external_connector--reference--group-001.md#canonical-e0a5ddb762b7db9a7a723f1bae6bf0f88bf21ade932a473e7c188fdcc11bf01e) |
| `ipsec.ike_parameters.ike_phase2_profile.name` | [ipsec.ike_parameters.ike_phase2_profile.name](data-sources--external_connector--reference--group-001.md#canonical-bb12dd429e7c1c5bfeb264c39a4d552a02963c3c4aa569ddeb9694779b5e8c76) |
| `ipsec.ike_parameters.ike_phase2_profile.namespace` | [ipsec.ike_parameters.ike_phase2_profile.namespace](data-sources--external_connector--reference--group-001.md#canonical-94b0411f8a4c7d27e8891eed1d1f529d1af3503b04c428592cba842f509285e3) |
| `ipsec.ike_parameters.ike_phase2_profile.tenant` | [ipsec.ike_parameters.ike_phase2_profile.tenant](data-sources--external_connector--reference--group-001.md#canonical-7e6dbd05996f4cb6430de94b9eaaa9fa64ddecbcd0dc9553f36da197a37fcee7) |
| `ipsec.ike_parameters.initiator` | [ipsec.ike_parameters.initiator](data-sources--external_connector--reference--group-001.md#canonical-1be17b065fd56e38c94501e3dd3dc270e985ccbcb773bb3bc1812c80fd6ffc4e) |
| `ipsec.ike_parameters.responder` | [ipsec.ike_parameters.responder](data-sources--external_connector--reference--group-001.md#canonical-54358cafea87668cdaeb8dc6aa0a0c00506a349829c0eadc824ed0ca0150e7ce) |
| `ipsec.ike_parameters.rm_hostname` | [ipsec.ike_parameters.rm_hostname](data-sources--external_connector--reference--group-001.md#canonical-17465660acf0bce6c8bac82393c58add98237b7488159c0900843cabfcb5389c) |
| `ipsec.ike_parameters.rm_ip_address` | [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-987892fadc0d7002b1f151dba7ad49396be1442302f33b73af39d5edcaddb1ce) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack` | [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-5623f789e018c7da428ce7a6a9b345f63e468bf190c535328fb2d6ed6577514d) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](data-sources--external_connector--reference--group-001.md#canonical-d4f2367a7e56389077bd7e4041abb302d4b66158c281af8e6f7a5ff1b19396f9) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr](data-sources--external_connector--reference--group-001.md#canonical-f9abd80b2c330e9537528cfc72ac46720bfd59730567b473f76c3b5f8cf58ebc) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](data-sources--external_connector--reference--group-001.md#canonical-9d69843f55ad45b006958d98c88b3855da5d21d1ab8c3ae41f54b69afe1423a1) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr](data-sources--external_connector--reference--group-001.md#canonical-700cceef8d845ec5f6fa9eea4eb7a6ae4235ad920e68881fb6f171139041c08f) |
| `ipsec.ike_parameters.rm_ip_address.ipv4` | [ipsec.ike_parameters.rm_ip_address.ipv4](data-sources--external_connector--reference--group-001.md#canonical-be1ab22294c448adb42367ae2144705dba4b4298ef2998308e5e0e4696dc9160) |
| `ipsec.ike_parameters.rm_ip_address.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.ipv4.addr](data-sources--external_connector--reference--group-001.md#canonical-af0e10508858d8de570bc3c2e2b90a76d7962c42e1ea3441612daec0efd4d581) |
| `ipsec.ike_parameters.rm_ip_address.ipv6` | [ipsec.ike_parameters.rm_ip_address.ipv6](data-sources--external_connector--reference--group-001.md#canonical-0ddd2ca3745d1becc7185e32c29dc3786129f3aa5310b30b1f88f5597a66a84d) |
| `ipsec.ike_parameters.rm_ip_address.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.ipv6.addr](data-sources--external_connector--reference--group-001.md#canonical-1f5ee8dceb03263d6fb8973b17ebd9ed26e56748ae0c42d46e07ac1b76b2551f) |
| `ipsec.ike_parameters.use_default_local_ike_id` | [ipsec.ike_parameters.use_default_local_ike_id](data-sources--external_connector--reference--group-001.md#canonical-c49989c77fe969d02556b7ff5699fa0618f3d3f7b2d826d350463858d97e17c2) |
| `ipsec.ike_parameters.use_default_remote_ike_id` | [ipsec.ike_parameters.use_default_remote_ike_id](data-sources--external_connector--reference--group-001.md#canonical-8094a403e1c207ec8488948d23b87a9825dd372b4e40e6a848c1b0087808f568) |
| `ipsec.ipsec_tunnel_parameters` | [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1fce779ff8074f01c4ddc9818d118ab2cf3f021b87ad336a300297feaee02a7f) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address` | [ipsec.ipsec_tunnel_parameters.peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-825b76b6e0835b14aa230ca28504a20bc18a922813d614773395c6f7ad830632) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` | [ipsec.ipsec_tunnel_parameters.peer_ip_address.addr](data-sources--external_connector--reference--group-001.md#canonical-24acbec4fd7cfb45a3cf71c7f8d6ca037b1373da05e53285f6e3d5be6a32cb4b) |
| `ipsec.ipsec_tunnel_parameters.psk` | [ipsec.ipsec_tunnel_parameters.psk](data-sources--external_connector--reference--group-001.md#canonical-4dac01432518c3bc65c43ae4fe42e4adcc486aaceefff830783547b1e6438836) |
| `ipsec.ipsec_tunnel_parameters.segment` | [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-4fb8d95153791ed359d3ee024407f3e7ef852c87701c6b28d846945b21cb3ab1) |
| `ipsec.ipsec_tunnel_parameters.segment.refs` | [ipsec.ipsec_tunnel_parameters.segment.refs](data-sources--external_connector--reference--group-001.md#canonical-679c6e4a8a9c65d610c418cc7d8e6172992ffee866272431d88be3fd84720bd2) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.kind` | [ipsec.ipsec_tunnel_parameters.segment.refs.kind](data-sources--external_connector--reference--group-001.md#canonical-84ff2371dd5b43e9a51b178b162dce8c96d2988b3b1d4934f9d8c8c93aae3d8d) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.name` | [ipsec.ipsec_tunnel_parameters.segment.refs.name](data-sources--external_connector--reference--group-001.md#canonical-658b9c6642d72b0f2d4338d1c02e5b2d1b37cf8c24c79d3886c38974936bb26c) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` | [ipsec.ipsec_tunnel_parameters.segment.refs.namespace](data-sources--external_connector--reference--group-001.md#canonical-24837b107b687333c9f2d12f163a0ca090f5efd922a80302e7fbef0b78ef513d) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` | [ipsec.ipsec_tunnel_parameters.segment.refs.tenant](data-sources--external_connector--reference--group-001.md#canonical-90a6fe70b90f8f359f1e52f775793b7ee282e8bedf01b442de08a7b435450403) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.uid` | [ipsec.ipsec_tunnel_parameters.segment.refs.uid](data-sources--external_connector--reference--group-001.md#canonical-0677a704b3476e7ac531fe7c5be98fc36e62321568ad1ee77078342f683e7497) |
| `ipsec.ipsec_tunnel_parameters.site_local_inside_network` | [ipsec.ipsec_tunnel_parameters.site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-6bd62f47ea4df47c47596ff98b5825fee0b0ff2f4d6ab03e4ebeab834ff6f2f1) |
| `ipsec.ipsec_tunnel_parameters.site_local_network` | [ipsec.ipsec_tunnel_parameters.site_local_network](data-sources--external_connector--reference--group-001.md#canonical-7e443230c5c59923030e6e90c8b786e71c3366556d9bc8e9036b417eccc67f80) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps` | [ipsec.ipsec_tunnel_parameters.tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-0c8eb65d0523bf9d6fb4a794857a9381462f5baf3ab2d6b472a00ecb14df1604) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.interface](data-sources--external_connector--reference--group-001.md#canonical-d97e6a134898c95f01c9ae63c4980fc2cb2e5409cff14fbbd0d2800cca9dbec2) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-63dedc161ffc6d347993f5d1e8b389e1eda821bcb33f97c03c0aef0aee7a89e1) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.node](data-sources--external_connector--reference--group-001.md#canonical-cbe4b937fc8a3b5c4c323498ce08d448bcc33db7f09a52f13a5a0f48372335b3) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-c86add8a3d1cec097f7412dcebd5cc318051dfb8b854c505eefb82fb438ad1b0) |
| `ipsec.ipsec_tunnel_parameters.tunnel_mtu` | [ipsec.ipsec_tunnel_parameters.tunnel_mtu](data-sources--external_connector--reference--group-001.md#canonical-2a002c06af4f0f8efc068269a7d4b57f82d43114db7e67cd273eecba65bf2136) |
| `labels` | [labels](data-sources--external_connector--reference--group-001.md#canonical-5ca2721972ae70a29563e0be8e77530d4b1c4c49be18e41235f101b941bf15e1) |
| `name` | [name](data-sources--external_connector--reference--group-001.md#canonical-12d478a5288698d9a4257f8bfdbe28e535dfb25506fd57220ca7860a03e8bd65) |
| `namespace` | [namespace](data-sources--external_connector--reference--group-001.md#canonical-c1353d1649d4f90e20f3acb61737378ab6161d78efccf24984aef715ff6d97d4) |

<a id="canonical-38dcd124a57f7726b1c72f6d85914aab7bfaf974527ba352e552a75a70249dde"></a>

## Next pages — Property reference / 9ca2569bd122 / 11

- [ce_site_reference](data-sources--external_connector--reference--group-001.md#canonical-09dea781a39324e59d78d3c5b6229f1e4279f242425a36e80e87b6d8373ec61c)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-09dea781a39324e59d78d3c5b6229f1e4279f242425a36e80e87b6d8373ec61c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cde386a3326080f612bf78b7c395c9eff9365e72dfb1eb2f457a8d263f3742e"></a>

## ce_site_reference — ce_site_reference / a2c055eb4167 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- ce_site_reference

<a id="canonical-912cefb2cae19844958aa9334db2fcf8d965d1736b82eaeee2e40404a4719a5c"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-29682121f01de50a8126c002967339ad1d119be77f9c241926bda75135962db8"></a>

## Direct properties — ce_site_reference / a2c055eb4167 / 3

<a id="canonical-e5eec23648ac135b9f796f42e8c7a8eee0b51c62e22cc913dcb61bd48ab10725"></a>

<a id="canonical-3ee60d05242d149046ccde64cec623948d76b0dafab568d4ceb80cbfe1720872"></a>

## name property — ce_site_reference / a2c055eb4167 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-da2df306b1419edea322844b645186e3d73f18c74eb0ced1ddd89e7b0e7e1253"></a>

<a id="canonical-a290bf382f7f49a8fc4ade5538ea415c61592a0b153bb75ce713ae84604de163"></a>

## namespace property — ce_site_reference / a2c055eb4167 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-c24bbdf9580d2a05b6fe74c0de9c128bf8df45aa43bff3ef6df8c2d1d9152b06"></a>

<a id="canonical-bef6325ceb5b63d8a72707d096f3f066615699998146be71754c9ce0f4195931"></a>

## tenant property — ce_site_reference / a2c055eb4167 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-e1f001bc680aed4b61ab1ef4b92745555dd0913ebd26bdf10c4c675d8c9c8e66"></a>

## Next pages — ce_site_reference / a2c055eb4167 / 7

- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdc879f17ba3b4f71960a907fae4e6076946bf8e45931e3078972cf9b30cc329"></a>

## gre — gre / d97d5d693230 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- gre

<a id="canonical-ff62476e01a590e20a11ed2b62b29f1d03c7483558dbd04d2c550e8b8ca9dc22"></a>

Type: `"single"`. Computed.

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

- [gre](data-sources--external_connector--reference--group-001.md#canonical-ff62476e01a590e20a11ed2b62b29f1d03c7483558dbd04d2c550e8b8ca9dc22)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-d09b756cc40d88ccfd47ab8d95dc4b012620da92abb968542c3335011a0e7861)

Select alternatives according to the provider validators above.

<a id="canonical-dcc644ed967ed2bc2c57062c4d51a6d394395ca5f2074b7c91bc085b165d36c0"></a>

## Direct properties — gre / d97d5d693230 / 3

- [gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f): complete subsection reference.

<a id="canonical-b687c3039bdb85e0ca63eb320c0e34b9e1fb2c615faca4429a93f5c6156484f7"></a>

## Next pages — gre / d97d5d693230 / 4

- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e4dc6678540911eb1378447ce70bddf6b34a43d3e4b92fb93bb81ee62803a39"></a>

## gre.gre_parameters — gre.gre_parameters / ace1b6f8ed37 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- gre.gre_parameters

<a id="canonical-e6575c113a22db1ca43504b05794881600e9998fb4f96d44e30f29cc6f650594"></a>

Type: `"single"`. Computed.

GRE configuration parameters required for GRE Connection type.

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

<a id="canonical-abd2d0e37bc03bc6a0b40c7987e22bd72a513b9a8aafd5dc935899acb5639040"></a>

## Direct properties — gre.gre_parameters / ace1b6f8ed37 / 3

- [peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-2e4a187e619055e61b7ce448401b9c2f1729cad9e94ebe50528650c998679aad): complete subsection reference.

- [segment](data-sources--external_connector--reference--group-001.md#canonical-c4ebed3e536bcd3bf79e88d87ef22e865658755183f56f808d7f0550878bac12): complete subsection reference.

- [site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-052a0deb22abfcc310c598bfe89d3285b817d1974faafc638a25eac3c8b5cc67): complete subsection reference.

- [site_local_network](data-sources--external_connector--reference--group-001.md#canonical-08a35002950c099be71d31131b389a5ecfee27c71798bb9e80e0e74fd12690b1): complete subsection reference.

- [tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-ba767a5c357e35bf412919074f6a2c72b121a5713bf327ae7966f484e0f9c3f0): complete subsection reference.

<a id="canonical-1233fdbb7a68df57b95f20e45234cb59eacc93e4db3ad945e4818a533a1a0513"></a>

<a id="canonical-500dd7135db4f249dcc2021e03ef48c93c11bb18db25fdc205de4a8118f7b21c"></a>

## tunnel_mtu property — gre.gre_parameters / ace1b6f8ed37 / 4

Type: `"number"`. Computed.

Configure MTU for the GRE tunnel interface.

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

<a id="canonical-4f9417c0d008515bfdf324d256d34a83aa10a76ede129563005a076f20777b95"></a>

## Next pages — gre.gre_parameters / ace1b6f8ed37 / 5

- [gre.gre_parameters.peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-2e4a187e619055e61b7ce448401b9c2f1729cad9e94ebe50528650c998679aad)
- [gre.gre_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-c4ebed3e536bcd3bf79e88d87ef22e865658755183f56f808d7f0550878bac12)
- [gre.gre_parameters.site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-052a0deb22abfcc310c598bfe89d3285b817d1974faafc638a25eac3c8b5cc67)
- [gre.gre_parameters.site_local_network](data-sources--external_connector--reference--group-001.md#canonical-08a35002950c099be71d31131b389a5ecfee27c71798bb9e80e0e74fd12690b1)
- [gre.gre_parameters.tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-ba767a5c357e35bf412919074f6a2c72b121a5713bf327ae7966f484e0f9c3f0)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-2e4a187e619055e61b7ce448401b9c2f1729cad9e94ebe50528650c998679aad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09611b5691b58456111f0f096b53c382451c4f70e5606dd4dab55649eeb5c2d6"></a>

## gre.gre_parameters.peer_ip_address — gre.gre_parameters.peer_ip_address / 7ea4a2a5dcb0 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- gre.gre_parameters.peer_ip_address

<a id="canonical-55a976bc510e3113f94c59cf30decffeaaa7ce8ca235de5662252c06e7b59ab6"></a>

Type: `"single"`. Computed.

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

<a id="canonical-da648a2be9500164417a6a8e581baa7e2e0edcfbb8985d4bb51a4b3c3e388a61"></a>

## Direct properties — gre.gre_parameters.peer_ip_address / 7ea4a2a5dcb0 / 3

<a id="canonical-705727c5659cfdf62078d15f0dfa590660879f89bdd9f591b6e848f4b22316bd"></a>

<a id="canonical-dbb103444fe81d4d409e48e55065256240145e9a1bc908620561b58a2db01a17"></a>

## addr property — gre.gre_parameters.peer_ip_address / 7ea4a2a5dcb0 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-f8ed8c204b92b3b9d18bb01e31d40681d711d617c37ed91b6339c1b3fb9b8576"></a>

## Next pages — gre.gre_parameters.peer_ip_address / 7ea4a2a5dcb0 / 5

- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-c4ebed3e536bcd3bf79e88d87ef22e865658755183f56f808d7f0550878bac12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-330994002f54c13de48172c1a2430cfde944205fc457b5e84c56ad415494cb24"></a>

## gre.gre_parameters.segment — gre.gre_parameters.segment / b853940e2ba1 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- gre.gre_parameters.segment

<a id="canonical-dfb5d9b018095a4876d25bfed7a9eb2bcd2b4befccc45df5507c0d8f39fa4f67"></a>

Type: `"single"`. Computed.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

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

<a id="canonical-e66255dd69ec8e95ce83ea86e049aaa7da648ab9d4f77950f21c62ed3f7104c5"></a>

## Direct properties — gre.gre_parameters.segment / b853940e2ba1 / 3

- [refs](data-sources--external_connector--reference--group-001.md#canonical-339bd2e60c6866ae835ead45fc5a575cd6d5104f5c4b6374d6d94e7b27c265ab): complete subsection reference.

<a id="canonical-b7eca93b093919026e6efb5c7f77898daaefa32a49b51382ade2f4564190b27d"></a>

## Next pages — gre.gre_parameters.segment / b853940e2ba1 / 4

- [gre.gre_parameters.segment.refs](data-sources--external_connector--reference--group-001.md#canonical-339bd2e60c6866ae835ead45fc5a575cd6d5104f5c4b6374d6d94e7b27c265ab)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-339bd2e60c6866ae835ead45fc5a575cd6d5104f5c4b6374d6d94e7b27c265ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e900def31fd12442e4744ae063d381d7adaf53d14997b688c6247a3e2e632e6"></a>

## gre.gre_parameters.segment.refs — gre.gre_parameters.segment.refs / d0d12157c083 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- [gre.gre_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-c4ebed3e536bcd3bf79e88d87ef22e865658755183f56f808d7f0550878bac12)
- gre.gre_parameters.segment.refs

<a id="canonical-3f7c4d99475b45884c182ab6165a424107d867a40e3a4e7c879eca4898f82d10"></a>

Type: `"list"`. Computed.

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

<a id="canonical-51f0be204787b6fa6ce2b811d67a360eedff7e1cc45f05f45d3bb2b6faba4699"></a>

## Direct properties — gre.gre_parameters.segment.refs / d0d12157c083 / 3

<a id="canonical-4d1c094170ad29a43cf23b4acb7e0fce0a7c182f936f6e45c6d32f04a49e2917"></a>

<a id="canonical-e82da8f109efb6c6aa253a1b163cb31adbae6281519615f93425d0b5b86c9ea7"></a>

## kind property — gre.gre_parameters.segment.refs / d0d12157c083 / 4

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

<a id="canonical-a1c4da369e713730fcaabda604e153c4cac406174cf3896a962378ac99a2f836"></a>

<a id="canonical-fa5f9208deb0ab4640d39a045e475559084688f686aaddf825a594dc8c141278"></a>

## name property — gre.gre_parameters.segment.refs / d0d12157c083 / 5

Type: `"string"`. Computed.

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

<a id="canonical-f03459834916bc458a99759444a4a0e50dfe9e67383f9870bb0bb27c38eb4e72"></a>

<a id="canonical-ce6cd37cfbeaa6cdf6063e2fec3957cc146ed761f9eb8cbf2c6dd7f4e0055de2"></a>

## namespace property — gre.gre_parameters.segment.refs / d0d12157c083 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-ee404db6709fa31d9a36a025686b2b0c3dd419f92756790589ed98f915dc9f6e"></a>

<a id="canonical-e7029dfac9736599655b9c39541b0b9a5c9a0b2f21c0c8ac5d4d3e57b83807da"></a>

## tenant property — gre.gre_parameters.segment.refs / d0d12157c083 / 7

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

<a id="canonical-740c4b39507e4d01d4da49dba6bbd87e3058fd52aaae8e2ca9e311a82b923ffe"></a>

<a id="canonical-ec0a66a15e73714e849eca566349ca86d2c1306632f4e1de8e042263b697028a"></a>

## uid property — gre.gre_parameters.segment.refs / d0d12157c083 / 8

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

<a id="canonical-48c057cf95f0fd18a7e487dcf3e04f0a9202d25dd10720af7c3d8ad0ecea7e65"></a>

## Next pages — gre.gre_parameters.segment.refs / d0d12157c083 / 9

- [gre.gre_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-c4ebed3e536bcd3bf79e88d87ef22e865658755183f56f808d7f0550878bac12)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-052a0deb22abfcc310c598bfe89d3285b817d1974faafc638a25eac3c8b5cc67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7471071222fa74efecb860acc49f164d137462131f4ae36b1c22828130fd9364"></a>

## gre.gre_parameters.site_local_inside_network — gre.gre_parameters.site_local_inside_network / c081a2a254c6 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- gre.gre_parameters.site_local_inside_network

<a id="canonical-38dea2993dfa5ea83b02fb50f2212464c01ace60d8877b83670e7acafb0f2572"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-4758cf8287010690f415f929bcb875731eefc98ce1da9e3558cc9a1856ab9d29"></a>

## Direct properties — gre.gre_parameters.site_local_inside_network / c081a2a254c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-923d14acc3dc7e055070598e93ee880f64171a9f86a93fc1f0a91dd29db0f920"></a>

## Next pages — gre.gre_parameters.site_local_inside_network / c081a2a254c6 / 4

- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-08a35002950c099be71d31131b389a5ecfee27c71798bb9e80e0e74fd12690b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1fbde9834fb2b57448964dad0d1cbd8a867b177703998a53c2cac38623c0b46"></a>

## gre.gre_parameters.site_local_network — gre.gre_parameters.site_local_network / a18c875c2ff6 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- gre.gre_parameters.site_local_network

<a id="canonical-2699bfb6f7577f7e1cb6166c490ca01026936e089ce3ff5c3941906a45eff276"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1cee3a6fee7ef5fbb84c240d674c1c26f146cc957ac63af1b94d07da744f36b9"></a>

## Direct properties — gre.gre_parameters.site_local_network / a18c875c2ff6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e057ac8bfe160edfce80d9f6fc7b0a1d9177673c30fb0eed57d4157712148eb6"></a>

## Next pages — gre.gre_parameters.site_local_network / a18c875c2ff6 / 4

- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-ba767a5c357e35bf412919074f6a2c72b121a5713bf327ae7966f484e0f9c3f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6300fb46827c111e1b96a34276317e3a4e66cd5c752ae70f5c74ab18812cefb9"></a>

## gre.gre_parameters.tunnel_eps — gre.gre_parameters.tunnel_eps / d03bfb95e24c / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-9f1781eca13d5f4f9f63c5ff61ec6ec419cd270306ef5f65a98b4a0df65a5949)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- gre.gre_parameters.tunnel_eps

<a id="canonical-28fd97bd92d93e588d69f3fabc410050207ede0f2e1807cb29f3d40af13e2b31"></a>

Type: `"list"`. Computed.

Configure tunnel parameters, source, destination, IP addresses.

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

<a id="canonical-69b7c2171676251ffac2a0fe820262b9357d29dc438c08f764003c0700f86c53"></a>

## Direct properties — gre.gre_parameters.tunnel_eps / d03bfb95e24c / 3

<a id="canonical-a65d91b3b23d800c45103b48c47c49c076db7ce85a1eb112306e7695eefda43f"></a>

<a id="canonical-318ea34949c4f7240d155a2f490df5f91eb689df6a34fd820c4b2295b7985f4c"></a>

## interface property — gre.gre_parameters.tunnel_eps / d03bfb95e24c / 4

Type: `"string"`. Computed.

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

<a id="canonical-1c8db96255256826719b3ed32cab586fc48497adbdc15464cb8c70df9390200c"></a>

<a id="canonical-040b2052057c96c8577d139ae860a823b1f3989c22a012e8d8d60366eb151b1b"></a>

## local_tunnel_ip property — gre.gre_parameters.tunnel_eps / d03bfb95e24c / 5

Type: `"string"`. Computed.

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

<a id="canonical-b4f29d8d16618ba7f1e708ba4b89baf848a5d62b3e348b86bfbdcee7d0bf1bf6"></a>

<a id="canonical-e669a024fee83af63a361eb539acb55e1bc13007ac78ecdc8d676c51a6bfa7e3"></a>

## node property — gre.gre_parameters.tunnel_eps / d03bfb95e24c / 6

Type: `"string"`. Computed.

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

<a id="canonical-02b1f127d72056e0cdebe847579e21f3e7803b6c8e1f93be43a3065e620d932f"></a>

<a id="canonical-8e28a07729c03cf8d39233c2fa2538b7ff4b563e94e9c457fba6cc6a063ce7ea"></a>

## remote_tunnel_ip property — gre.gre_parameters.tunnel_eps / d03bfb95e24c / 7

Type: `"string"`. Computed.

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

<a id="canonical-a634cdaf4cb42d25d2b75b00711216dd33e24d7d024cbe466ab7fc27caf44664"></a>

## Next pages — gre.gre_parameters.tunnel_eps / d03bfb95e24c / 8

- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-d544434b53111705492f0038273477003629823b90b7b7dfe70e130aef83fa2f)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94b999b6409c135a3431e5c77b7647b6c3287075a6cb98003f2977a4aa90aee4"></a>

## ipsec — ipsec / 137fee2e46e0 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- ipsec

<a id="canonical-d09b756cc40d88ccfd47ab8d95dc4b012620da92abb968542c3335011a0e7861"></a>

Type: `"single"`. Computed.

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

<a id="canonical-a2e42ef4082e2f8877236dceb2b252389317552721d973d3e0c730918051ea2d"></a>

## Direct properties — ipsec / 137fee2e46e0 / 3

- [ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b): complete subsection reference.

- [ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf): complete subsection reference.

<a id="canonical-4643d600691adbadcb435d93091ae512a98784fda7711b5b328812ffdb28fa25"></a>

## Next pages — ipsec / 137fee2e46e0 / 4

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f49e98dbb9816f06db4d5e78eb383cd4c852034fdffef0b24ecea6161b1995a0"></a>

## ipsec.ike_parameters — ipsec.ike_parameters / b6ad07499869 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- ipsec.ike_parameters

<a id="canonical-8d6fa631584e216a8a3f68044ebc83a95e67455f512b0448adaa3f2d760e40f2"></a>

Type: `"single"`. Computed.

IKE configuration parameters required for IPsec Connection type.

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

<a id="canonical-8cf2dcd4b97ac5e79efa6a3a0f587b177951ccbf7b43277fe34fc79a13a977b7"></a>

## Direct properties — ipsec.ike_parameters / b6ad07499869 / 3

- [dpd_disabled](data-sources--external_connector--reference--group-001.md#canonical-6d1fecff641a3b01332a787938278014a03b842166061de7b441403a05cf0aa2): complete subsection reference.

- [dpd_keep_alive_timer](data-sources--external_connector--reference--group-001.md#canonical-9a87a4f7849369153e9a76917bb2b42172a8a315c165865d697067ed82428416): complete subsection reference.

- [ike_phase1_profile](data-sources--external_connector--reference--group-001.md#canonical-c70e61de8a36295cf574079acef06ef4ff493484f28b74c4349825f39bd22262): complete subsection reference.

- [ike_phase2_profile](data-sources--external_connector--reference--group-001.md#canonical-a7c68acb6b74d8350009aca51280d4a7b89063d9479dbe1aeb74955deb5c4188): complete subsection reference.

- [initiator](data-sources--external_connector--reference--group-001.md#canonical-0944fbcf0598e4c2da8aa95d0434f79292ec34b491d0b6d3f53f6dd66be3abb2): complete subsection reference.

- [responder](data-sources--external_connector--reference--group-001.md#canonical-7a76e8b19daac49d60d8c65b7ad516f1644059154121e743f1fb134dc24d0f24): complete subsection reference.

<a id="canonical-17465660acf0bce6c8bac82393c58add98237b7488159c0900843cabfcb5389c"></a>

<a id="canonical-3d3a3e5bc4525215d6281bf488cb6fc89b78a680f0526cf65447e4c304ea50b2"></a>

## rm_hostname property — ipsec.ike_parameters / b6ad07499869 / 4

Type: `"string"`. Computed.

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

- [rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6): complete subsection reference.

- [use_default_local_ike_id](data-sources--external_connector--reference--group-001.md#canonical-272471551ce013550118bd257637d6ed7a0ece49e1904ac8f3d63acf55202bb8): complete subsection reference.

- [use_default_remote_ike_id](data-sources--external_connector--reference--group-001.md#canonical-58aa95a4b691ba9c336c76f134e734dda491b99d7310649bdc01d3a2e24d1c38): complete subsection reference.

<a id="canonical-a73fedbfddcaf36ed5c9fcefd1ae5b171f673716a57c1e6a38b84c6f2e0e9fe2"></a>

## Next pages — ipsec.ike_parameters / b6ad07499869 / 5

- [ipsec.ike_parameters.dpd_disabled](data-sources--external_connector--reference--group-001.md#canonical-6d1fecff641a3b01332a787938278014a03b842166061de7b441403a05cf0aa2)
- [ipsec.ike_parameters.dpd_keep_alive_timer](data-sources--external_connector--reference--group-001.md#canonical-9a87a4f7849369153e9a76917bb2b42172a8a315c165865d697067ed82428416)
- [ipsec.ike_parameters.ike_phase1_profile](data-sources--external_connector--reference--group-001.md#canonical-c70e61de8a36295cf574079acef06ef4ff493484f28b74c4349825f39bd22262)
- [ipsec.ike_parameters.ike_phase2_profile](data-sources--external_connector--reference--group-001.md#canonical-a7c68acb6b74d8350009aca51280d4a7b89063d9479dbe1aeb74955deb5c4188)
- [ipsec.ike_parameters.initiator](data-sources--external_connector--reference--group-001.md#canonical-0944fbcf0598e4c2da8aa95d0434f79292ec34b491d0b6d3f53f6dd66be3abb2)
- [ipsec.ike_parameters.responder](data-sources--external_connector--reference--group-001.md#canonical-7a76e8b19daac49d60d8c65b7ad516f1644059154121e743f1fb134dc24d0f24)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- [ipsec.ike_parameters.use_default_local_ike_id](data-sources--external_connector--reference--group-001.md#canonical-272471551ce013550118bd257637d6ed7a0ece49e1904ac8f3d63acf55202bb8)
- [ipsec.ike_parameters.use_default_remote_ike_id](data-sources--external_connector--reference--group-001.md#canonical-58aa95a4b691ba9c336c76f134e734dda491b99d7310649bdc01d3a2e24d1c38)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-6d1fecff641a3b01332a787938278014a03b842166061de7b441403a05cf0aa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d7a1600fd1e3e9eff8e85f36357293cc903209aff1b4aa7685ac474f2b42d6e"></a>

## ipsec.ike_parameters.dpd_disabled — ipsec.ike_parameters.dpd_disabled / 14ed04fc938c / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.dpd_disabled

<a id="canonical-13958225e44f15c2744f7708aad30f329c867acdf6a6f594f5df74925ee7ccd4"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-66043a04ed53dfd5a9f9d9014837b50677e463b3365752f310d58dfdce0648a5"></a>

## Direct properties — ipsec.ike_parameters.dpd_disabled / 14ed04fc938c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c4fd4e0d1c89d8d6b874e90ab3f9c80582cb7850afdfb5910714404d09f423a"></a>

## Next pages — ipsec.ike_parameters.dpd_disabled / 14ed04fc938c / 4

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-9a87a4f7849369153e9a76917bb2b42172a8a315c165865d697067ed82428416"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e249dc9d271bb814ce7c93c25d22879782842ee77ea2ae84252ace4a400c20e"></a>

## ipsec.ike_parameters.dpd_keep_alive_timer — ipsec.ike_parameters.dpd_keep_alive_timer / 9540d1fb7e8a / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.dpd_keep_alive_timer

<a id="canonical-3e26b634a9c76d9266128338532af832c486f043fcd8e9eca221b3a3d9309d87"></a>

Type: `"single"`. Computed.

Configuration parameter for dpd keep alive timer.

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

<a id="canonical-9e85a579ee48716cb0435c10840c392bd25749caa19cc1f94f0c47c852026c3f"></a>

## Direct properties — ipsec.ike_parameters.dpd_keep_alive_timer / 9540d1fb7e8a / 3

<a id="canonical-1ab0e6975e186637d91682861180a50f5b3d1ed57eaf3297cfe0668df47b2ecc"></a>

<a id="canonical-93ba7d23ae40063fb7a06f4047250e4ec7fdeff7f339f60e8145e44f9b961156"></a>

## timeout property — ipsec.ike_parameters.dpd_keep_alive_timer / 9540d1fb7e8a / 4

Type: `"number"`. Computed.

Keepalive Timer. Operation timeout duration

Upstream description:

Operation timeout duration

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

<a id="canonical-e94995b646f6ea018d3933f12d1131f2c8206407ec2ec5fe582118b5681b9f58"></a>

## Next pages — ipsec.ike_parameters.dpd_keep_alive_timer / 9540d1fb7e8a / 5

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-c70e61de8a36295cf574079acef06ef4ff493484f28b74c4349825f39bd22262"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1af0864e5d7a05e6962291ddc844b40f1870934ef97b631328a60c67962f534e"></a>

## ipsec.ike_parameters.ike_phase1_profile — ipsec.ike_parameters.ike_phase1_profile / 3834c7456c79 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.ike_phase1_profile

<a id="canonical-1ed4c8d0f51d9b258c3810520833d4e3a9682bf923a70191e33431da7fb5b423"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-024e5f772dc296ebe50ca5c8d83cfc611aff54ed1a96d90749d36a485fcc84d2"></a>

## Direct properties — ipsec.ike_parameters.ike_phase1_profile / 3834c7456c79 / 3

<a id="canonical-03ef8cd99768d1875b4dace0973fd65d74bd14883c8ee75bc7c27b845b34d67e"></a>

<a id="canonical-01b5b7f21e2cd81c94bc74cb8898b00df00247b8d8b27fce129d841c839749a9"></a>

## name property — ipsec.ike_parameters.ike_phase1_profile / 3834c7456c79 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-46c4f25bf74b24e6fe20102ce1315fb32847e94ff6142f2307368bade2b35972"></a>

<a id="canonical-a47558362f7e3114de8a57795ce0ba2564cc19d04dd298dff4486b9e2a3aa56e"></a>

## namespace property — ipsec.ike_parameters.ike_phase1_profile / 3834c7456c79 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-f8a5b5e021a71c60ace9725113ece4f333e993353b6aa4760f78eb6110ad87a2"></a>

<a id="canonical-45d7841e785422219772d354b3c0d197e39596d2c9c126c6a971ef5f2bb85f60"></a>

## tenant property — ipsec.ike_parameters.ike_phase1_profile / 3834c7456c79 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-e850d3c2d0cceed0180ab5ff0c3dee25e11b0e66aa614b999b8f15787625de5d"></a>

## Next pages — ipsec.ike_parameters.ike_phase1_profile / 3834c7456c79 / 7

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-a7c68acb6b74d8350009aca51280d4a7b89063d9479dbe1aeb74955deb5c4188"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1cd1d833d978842c955476a1c1ba65bbc17300779876ddb052998b328114bd3"></a>

## ipsec.ike_parameters.ike_phase2_profile — ipsec.ike_parameters.ike_phase2_profile / 0e901ca9bf4f / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.ike_phase2_profile

<a id="canonical-e0a5ddb762b7db9a7a723f1bae6bf0f88bf21ade932a473e7c188fdcc11bf01e"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-e03f54f971a8968ede6422baa4a153a57920d1a652887e0ec5af210ca26b8626"></a>

## Direct properties — ipsec.ike_parameters.ike_phase2_profile / 0e901ca9bf4f / 3

<a id="canonical-bb12dd429e7c1c5bfeb264c39a4d552a02963c3c4aa569ddeb9694779b5e8c76"></a>

<a id="canonical-2a2763b5c411c110f0d0957d9eb3e8d1fc648e19316021fa8f27f50234f4cd96"></a>

## name property — ipsec.ike_parameters.ike_phase2_profile / 0e901ca9bf4f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-94b0411f8a4c7d27e8891eed1d1f529d1af3503b04c428592cba842f509285e3"></a>

<a id="canonical-0e55982276fc296733c47575b68566df43e3ae1644259253534fc3a914b7582a"></a>

## namespace property — ipsec.ike_parameters.ike_phase2_profile / 0e901ca9bf4f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-7e6dbd05996f4cb6430de94b9eaaa9fa64ddecbcd0dc9553f36da197a37fcee7"></a>

<a id="canonical-180f70a418230bb3d3e8bbb555b53ad3d63412bf5e2212f4189ec4fcd34bd606"></a>

## tenant property — ipsec.ike_parameters.ike_phase2_profile / 0e901ca9bf4f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-8867ad68712ada3e6f31ee2486b2bfc7416fea91c58bbe2235279d0d3936acb8"></a>

## Next pages — ipsec.ike_parameters.ike_phase2_profile / 0e901ca9bf4f / 7

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-0944fbcf0598e4c2da8aa95d0434f79292ec34b491d0b6d3f53f6dd66be3abb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2afdff43f0a8cdf96b4832c85fc14bc2417a3bf43d3140e41c4b5c5859ae8612"></a>

## ipsec.ike_parameters.initiator — ipsec.ike_parameters.initiator / fcbd8382fa09 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.initiator

<a id="canonical-1be17b065fd56e38c94501e3dd3dc270e985ccbcb773bb3bc1812c80fd6ffc4e"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ff63a73c11248636eac98f15edcf293fb941cc46859e5b0b70c2ca490ab6471e"></a>

## Direct properties — ipsec.ike_parameters.initiator / fcbd8382fa09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14f9c0caa793a744c5bf3c51fd6bae8479f114ea35e544e0513459c78bc3e01b"></a>

## Next pages — ipsec.ike_parameters.initiator / fcbd8382fa09 / 4

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-7a76e8b19daac49d60d8c65b7ad516f1644059154121e743f1fb134dc24d0f24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e0552709e703ada4e680a57198e5bda0a32f1cc848bf47ae2ac478287fe4328"></a>

## ipsec.ike_parameters.responder — ipsec.ike_parameters.responder / 2a3171b7bfb7 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.responder

<a id="canonical-54358cafea87668cdaeb8dc6aa0a0c00506a349829c0eadc824ed0ca0150e7ce"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b81ece91644856fc9bf59a81df9636127b67ef1fd36a9451705ad2e4208e3672"></a>

## Direct properties — ipsec.ike_parameters.responder / 2a3171b7bfb7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba54c0e64a4261fca617ed2df299ab62113182562fe94076fe52ba814479acb0"></a>

## Next pages — ipsec.ike_parameters.responder / 2a3171b7bfb7 / 4

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68f28e9152bc3d00d3c1eb47f23332815b1afbd679fbcb902fa26f11f33a2cfc"></a>

## ipsec.ike_parameters.rm_ip_address — ipsec.ike_parameters.rm_ip_address / 6e17ad65365e / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.rm_ip_address

<a id="canonical-987892fadc0d7002b1f151dba7ad49396be1442302f33b73af39d5edcaddb1ce"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

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

<a id="canonical-a5710773eea7012ed8c64cd3a42131218a2da4349c430bac1567da7e72793e7e"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address / 6e17ad65365e / 3

- [dual_stack](data-sources--external_connector--reference--group-001.md#canonical-3649db33994342501370348d582b47ab60cc7a32d4fdb4629859aad69fb5a7d8): complete subsection reference.

- [ipv4](data-sources--external_connector--reference--group-001.md#canonical-e088a8a0a899c17bedbfbb4d8d6d23ce6230b36208e01433a18845c6fc5c5088): complete subsection reference.

- [ipv6](data-sources--external_connector--reference--group-001.md#canonical-c7858b8917b4955035253a5c850050f2ca6efd98344e6668a67c5ae7d234722d): complete subsection reference.

<a id="canonical-4cdd8d64061b1f509939c7707e30141df26cc2b701ee135fe65ee31b711d77d6"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address / 6e17ad65365e / 4

- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-3649db33994342501370348d582b47ab60cc7a32d4fdb4629859aad69fb5a7d8)
- [ipsec.ike_parameters.rm_ip_address.ipv4](data-sources--external_connector--reference--group-001.md#canonical-e088a8a0a899c17bedbfbb4d8d6d23ce6230b36208e01433a18845c6fc5c5088)
- [ipsec.ike_parameters.rm_ip_address.ipv6](data-sources--external_connector--reference--group-001.md#canonical-c7858b8917b4955035253a5c850050f2ca6efd98344e6668a67c5ae7d234722d)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-3649db33994342501370348d582b47ab60cc7a32d4fdb4629859aad69fb5a7d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8449a791c090f4dd18946c220db9dc6521d19e4a7e983b9630b8762b09598ae9"></a>

## ipsec.ike_parameters.rm_ip_address.dual_stack — ipsec.ike_parameters.rm_ip_address.dual_stack / 9ceafd485c95 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- ipsec.ike_parameters.rm_ip_address.dual_stack

<a id="canonical-5623f789e018c7da428ce7a6a9b345f63e468bf190c535328fb2d6ed6577514d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-e8529fd265d7ac572863e33b4e8b8b08a3bb5d97137765b8b233ee112c56adc8"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.dual_stack / 9ceafd485c95 / 3

- [ipv4](data-sources--external_connector--reference--group-001.md#canonical-43e6bd79213b7d58184f9477585835ecf023ebc4151636c043b750dc3802f62d): complete subsection reference.

- [ipv6](data-sources--external_connector--reference--group-001.md#canonical-4a7ffecc886e7807d9cf166baed6438d4ad9163de0cbb6b38fa26855ae94e751): complete subsection reference.

<a id="canonical-4a464636a86a3d5b74b236a4fb98d8884e9aef8a67b8f1b4df653b1d03ee2c86"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.dual_stack / 9ceafd485c95 / 4

- [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](data-sources--external_connector--reference--group-001.md#canonical-43e6bd79213b7d58184f9477585835ecf023ebc4151636c043b750dc3802f62d)
- [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](data-sources--external_connector--reference--group-001.md#canonical-4a7ffecc886e7807d9cf166baed6438d4ad9163de0cbb6b38fa26855ae94e751)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-43e6bd79213b7d58184f9477585835ecf023ebc4151636c043b750dc3802f62d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d131a6d14d70b4885b09ab2ae53652c3422bb63618adb5a8c6505e25b9d01ff"></a>

## ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 461cc1ceae28 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-3649db33994342501370348d582b47ab60cc7a32d4fdb4629859aad69fb5a7d8)
- ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4

<a id="canonical-d4f2367a7e56389077bd7e4041abb302d4b66158c281af8e6f7a5ff1b19396f9"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2be4c5cc1fec13994ec30d1a5f7021240a738db8be5ad2270b73aa2e1271e926"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 461cc1ceae28 / 3

<a id="canonical-f9abd80b2c330e9537528cfc72ac46720bfd59730567b473f76c3b5f8cf58ebc"></a>

<a id="canonical-f00ffb48abcb99fa7de326ce2d2a05cf68d7c24093e009cba57d816dd0d551a3"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 461cc1ceae28 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-9c6b11d8f70902717784ae243ee2560ecaccbea3a8f5ba18b0e2a53a96935202"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4 / 461cc1ceae28 / 5

- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-3649db33994342501370348d582b47ab60cc7a32d4fdb4629859aad69fb5a7d8)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-4a7ffecc886e7807d9cf166baed6438d4ad9163de0cbb6b38fa26855ae94e751"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c0ba4ec074a9e6aaf819f5deaabaa874887248124d927b267221f1baa935125"></a>

## ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 1b6206c9f138 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-3649db33994342501370348d582b47ab60cc7a32d4fdb4629859aad69fb5a7d8)
- ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6

<a id="canonical-9d69843f55ad45b006958d98c88b3855da5d21d1ab8c3ae41f54b69afe1423a1"></a>

Type: `"single"`. Computed.

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

<a id="canonical-d658c9542dc1d43345ae7f27673668b8efd38ce3449a76f4713c468466071a0e"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 1b6206c9f138 / 3

<a id="canonical-700cceef8d845ec5f6fa9eea4eb7a6ae4235ad920e68881fb6f171139041c08f"></a>

<a id="canonical-183d86872b0178c2d29cd74aaa324e705aed10b1d12893518a13dbe8a2a1027f"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 1b6206c9f138 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-166edf2333aee07a4344a43e6f08199d24f1d8efacdcf144776d00baa854c0da"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 / 1b6206c9f138 / 5

- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-3649db33994342501370348d582b47ab60cc7a32d4fdb4629859aad69fb5a7d8)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-e088a8a0a899c17bedbfbb4d8d6d23ce6230b36208e01433a18845c6fc5c5088"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1275a241a732a53b3030d613e79150157cfa79c81d1506b737d0832f1096dd41"></a>

## ipsec.ike_parameters.rm_ip_address.ipv4 — ipsec.ike_parameters.rm_ip_address.ipv4 / 401054068500 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- ipsec.ike_parameters.rm_ip_address.ipv4

<a id="canonical-be1ab22294c448adb42367ae2144705dba4b4298ef2998308e5e0e4696dc9160"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c3602397ca20e587d33d450a893f588bbd37a21298bc65da7f1b335e49b9d017"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.ipv4 / 401054068500 / 3

<a id="canonical-af0e10508858d8de570bc3c2e2b90a76d7962c42e1ea3441612daec0efd4d581"></a>

<a id="canonical-1e3e12b6cc6222524828647e05618b3e8ee360b9db34d53ad576a06f9a4df9c5"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.ipv4 / 401054068500 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-13f855c8173d018876aaa6e3be6752d2dff095c2384d980f6e71d38273de2d15"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.ipv4 / 401054068500 / 5

- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-c7858b8917b4955035253a5c850050f2ca6efd98344e6668a67c5ae7d234722d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1324865d8135813a61619c7c5c6dc307b3ea749a622faf731542dfdce84a4d5f"></a>

## ipsec.ike_parameters.rm_ip_address.ipv6 — ipsec.ike_parameters.rm_ip_address.ipv6 / c8de0b7253dd / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- ipsec.ike_parameters.rm_ip_address.ipv6

<a id="canonical-0ddd2ca3745d1becc7185e32c29dc3786129f3aa5310b30b1f88f5597a66a84d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-4edc6b0f800e7494b4ff8547da12b4c45ed70914c201ea2bab622ff3a73739a4"></a>

## Direct properties — ipsec.ike_parameters.rm_ip_address.ipv6 / c8de0b7253dd / 3

<a id="canonical-1f5ee8dceb03263d6fb8973b17ebd9ed26e56748ae0c42d46e07ac1b76b2551f"></a>

<a id="canonical-ae77d2aab91e583067ff118d4c6836420d8fd4f7ae37f762f68a6680a89e1db1"></a>

## addr property — ipsec.ike_parameters.rm_ip_address.ipv6 / c8de0b7253dd / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-ad753b5502af83123d06b4a8a1a6429b6a3b855ecc88037eb8c5fa1bc806e534"></a>

## Next pages — ipsec.ike_parameters.rm_ip_address.ipv6 / c8de0b7253dd / 5

- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-78f39410e3b8536eb3af2311078686d53650673b8d74b810fe260aba0ed18ff6)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-272471551ce013550118bd257637d6ed7a0ece49e1904ac8f3d63acf55202bb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f56c2b88f38ab252b9c1d91b69914475fe6f3ffe19b373fc5f48a10df227cc9e"></a>

## ipsec.ike_parameters.use_default_local_ike_id — ipsec.ike_parameters.use_default_local_ike_id / 2c5441bd2ce4 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.use_default_local_ike_id

<a id="canonical-c49989c77fe969d02556b7ff5699fa0618f3d3f7b2d826d350463858d97e17c2"></a>

Type: `"single"`. Computed.

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

<a id="canonical-366c711081773347997983e0abf3a3cbe54b4addd5aeee7030535c132d4a298d"></a>

## Direct properties — ipsec.ike_parameters.use_default_local_ike_id / 2c5441bd2ce4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d0e0237c23a7d6f426e2c73d7bec0d0025d2a78d303f357d2ae7ac060ee14e6"></a>

## Next pages — ipsec.ike_parameters.use_default_local_ike_id / 2c5441bd2ce4 / 4

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-58aa95a4b691ba9c336c76f134e734dda491b99d7310649bdc01d3a2e24d1c38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7aa495e01defd70e97971485e90ecdc75ad7cd2d64a14112e244a5b4fa43f7f"></a>

## ipsec.ike_parameters.use_default_remote_ike_id — ipsec.ike_parameters.use_default_remote_ike_id / 2a8991f34d35 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- ipsec.ike_parameters.use_default_remote_ike_id

<a id="canonical-8094a403e1c207ec8488948d23b87a9825dd372b4e40e6a848c1b0087808f568"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d9621465cb900a97abd19b2ff93df20cf513b868fba1d3009a4073e439c9e62a"></a>

## Direct properties — ipsec.ike_parameters.use_default_remote_ike_id / 2a8991f34d35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f0d037bd4a69acbf2bb84436e2e10e5b793aeffe0f320c32be4097a1dc982d5"></a>

## Next pages — ipsec.ike_parameters.use_default_remote_ike_id / 2a8991f34d35 / 4

- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-7c0af9da6fce39bd13de3b782bc7d684db977a5173d7afa79f6b5f250aad332b)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2183ae42c770f82f157af681181e66caaa45699e306b04b9c71d9122d9464255"></a>

## ipsec.ipsec_tunnel_parameters — ipsec.ipsec_tunnel_parameters / 84660467a301 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- ipsec.ipsec_tunnel_parameters

<a id="canonical-1fce779ff8074f01c4ddc9818d118ab2cf3f021b87ad336a300297feaee02a7f"></a>

Type: `"single"`. Computed.

In this section, we will configure the tunnel parameters, source, destination, IP addresses, and
segment.

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

<a id="canonical-76ad9c70687baa9e35d301d263d575e2a38f6f5dd9e8c00bace54ece411079b6"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters / 84660467a301 / 3

- [peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-88808e05ca17850b1654a5e61a047629220bf991471247dafe8d728659f98285): complete subsection reference.

<a id="canonical-4dac01432518c3bc65c43ae4fe42e4adcc486aaceefff830783547b1e6438836"></a>

<a id="canonical-f5b8217303582f11acfc93aac6dc2c7ad3dc28657f8ff7a2cbd1a8f46c371744"></a>

## psk property — ipsec.ipsec_tunnel_parameters / 84660467a301 / 4

Type: `"string"`. Computed.

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

- [segment](data-sources--external_connector--reference--group-001.md#canonical-4405808c5598812dc00f20d3b1dba6287ad6bb4def8875d92b951bbf6a79b6c1): complete subsection reference.

- [site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-ccc70534aaff58c990a5283b2a25ebffb64243bfc02c7c9108340eb2b3cc5eaf): complete subsection reference.

- [site_local_network](data-sources--external_connector--reference--group-001.md#canonical-04b993e915a8dad63fef9860ae39edfb47583a7e7de88b234c83682e02e1bd8e): complete subsection reference.

- [tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-63e1300761efb020dbe6ee0af48e1537fc77f867e4aabf6a8b3eaa34f31b4418): complete subsection reference.

<a id="canonical-2a002c06af4f0f8efc068269a7d4b57f82d43114db7e67cd273eecba65bf2136"></a>

<a id="canonical-cbb7ef0126e2640e8435fca343aa06f67ceb00190dd8b8542db0bd9bd15fc769"></a>

## tunnel_mtu property — ipsec.ipsec_tunnel_parameters / 84660467a301 / 5

Type: `"number"`. Computed.

The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without
needing to be fragmented.

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

<a id="canonical-c384a5314cbb9928f3df0f206d39aa256494b97427352956950b74939ed11e2a"></a>

## Next pages — ipsec.ipsec_tunnel_parameters / 84660467a301 / 6

- [ipsec.ipsec_tunnel_parameters.peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-88808e05ca17850b1654a5e61a047629220bf991471247dafe8d728659f98285)
- [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-4405808c5598812dc00f20d3b1dba6287ad6bb4def8875d92b951bbf6a79b6c1)
- [ipsec.ipsec_tunnel_parameters.site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-ccc70534aaff58c990a5283b2a25ebffb64243bfc02c7c9108340eb2b3cc5eaf)
- [ipsec.ipsec_tunnel_parameters.site_local_network](data-sources--external_connector--reference--group-001.md#canonical-04b993e915a8dad63fef9860ae39edfb47583a7e7de88b234c83682e02e1bd8e)
- [ipsec.ipsec_tunnel_parameters.tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-63e1300761efb020dbe6ee0af48e1537fc77f867e4aabf6a8b3eaa34f31b4418)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-88808e05ca17850b1654a5e61a047629220bf991471247dafe8d728659f98285"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8342205885e8001be8361d5a6f68a24460b9eb329dc998c02aea5140faff84b"></a>

## ipsec.ipsec_tunnel_parameters.peer_ip_address — ipsec.ipsec_tunnel_parameters.peer_ip_address / 176eeb854f29 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- ipsec.ipsec_tunnel_parameters.peer_ip_address

<a id="canonical-825b76b6e0835b14aa230ca28504a20bc18a922813d614773395c6f7ad830632"></a>

Type: `"single"`. Computed.

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

<a id="canonical-ea26704c3ad9735bd95c17a7d0627dc08fe28cdb9827c443e4373d12679091fa"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.peer_ip_address / 176eeb854f29 / 3

<a id="canonical-24acbec4fd7cfb45a3cf71c7f8d6ca037b1373da05e53285f6e3d5be6a32cb4b"></a>

<a id="canonical-e8023774aea03d9bca6c936e45fda4e1b45e29cad86c493a3a82f39255b6998f"></a>

## addr property — ipsec.ipsec_tunnel_parameters.peer_ip_address / 176eeb854f29 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-8c8fc30a32d5be7cde0c1c0e84a2b138b0096381c1802826e6439e7db9e9a62b"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.peer_ip_address / 176eeb854f29 / 5

- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-4405808c5598812dc00f20d3b1dba6287ad6bb4def8875d92b951bbf6a79b6c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-875613adb71cef12b2f806f68fcc4a7c543a849d56d4db9a86ad2b5a7ab9be01"></a>

## ipsec.ipsec_tunnel_parameters.segment — ipsec.ipsec_tunnel_parameters.segment / e3f022318df3 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- ipsec.ipsec_tunnel_parameters.segment

<a id="canonical-4fb8d95153791ed359d3ee024407f3e7ef852c87701c6b28d846945b21cb3ab1"></a>

Type: `"single"`. Computed.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

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

<a id="canonical-5ef7fb9f70d718131bb03dada18c74afccb182b932aaaf224386d95825b7e46d"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.segment / e3f022318df3 / 3

- [refs](data-sources--external_connector--reference--group-001.md#canonical-bf5fa4f30a7de57d894adfc3bcfac15c1adf18b54ece2ad595b779beefd48304): complete subsection reference.

<a id="canonical-6cdef8fe35d9efe89910e053716d72585c8d5e39ac9af8124136d051ca081a2f"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.segment / e3f022318df3 / 4

- [ipsec.ipsec_tunnel_parameters.segment.refs](data-sources--external_connector--reference--group-001.md#canonical-bf5fa4f30a7de57d894adfc3bcfac15c1adf18b54ece2ad595b779beefd48304)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-bf5fa4f30a7de57d894adfc3bcfac15c1adf18b54ece2ad595b779beefd48304"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6c81f71b90f91e50197715102e41f84195a8b8d236638ac5a90fb7045a096cb"></a>

## ipsec.ipsec_tunnel_parameters.segment.refs — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-4405808c5598812dc00f20d3b1dba6287ad6bb4def8875d92b951bbf6a79b6c1)
- ipsec.ipsec_tunnel_parameters.segment.refs

<a id="canonical-679c6e4a8a9c65d610c418cc7d8e6172992ffee866272431d88be3fd84720bd2"></a>

Type: `"list"`. Computed.

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

<a id="canonical-c31adfc9194126f0db55ab72c6c05958ba003d4f62974fe7e57e47ab5733c60e"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 3

<a id="canonical-84ff2371dd5b43e9a51b178b162dce8c96d2988b3b1d4934f9d8c8c93aae3d8d"></a>

<a id="canonical-80579d51cba49aa4c8f9d6ea02f4fb855410d38a5c735afb2a9fac0eb53fef67"></a>

## kind property — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 4

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

<a id="canonical-658b9c6642d72b0f2d4338d1c02e5b2d1b37cf8c24c79d3886c38974936bb26c"></a>

<a id="canonical-e14829f1178908a4377548902d05b6005498505a4e5d4d3f5c22f388bf2f45c9"></a>

## name property — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 5

Type: `"string"`. Computed.

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

<a id="canonical-24837b107b687333c9f2d12f163a0ca090f5efd922a80302e7fbef0b78ef513d"></a>

<a id="canonical-7e7a2e020b24167f92e2e80c8b78fb5494d44a0d6e1432a998980dc78884b6a6"></a>

## namespace property — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-90a6fe70b90f8f359f1e52f775793b7ee282e8bedf01b442de08a7b435450403"></a>

<a id="canonical-a1e940011742db1cd5c11b805c5585e7a90f431a8e0d00510f8736743a806a41"></a>

## tenant property — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 7

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

<a id="canonical-0677a704b3476e7ac531fe7c5be98fc36e62321568ad1ee77078342f683e7497"></a>

<a id="canonical-96a197afb88bab48dc6e03246e76bef315a9bf8d6959da562462711085b33f84"></a>

## uid property — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 8

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

<a id="canonical-49ac5f63befb59fa068df107314565d80fa1d6c93377fc4cf9ced4a73ca69d36"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.segment.refs / 676d7b78b558 / 9

- [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-4405808c5598812dc00f20d3b1dba6287ad6bb4def8875d92b951bbf6a79b6c1)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-ccc70534aaff58c990a5283b2a25ebffb64243bfc02c7c9108340eb2b3cc5eaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dccb230ac33424a3de09981a9bf4513e6e6420c8fef7ebb3c3c4770f32e2c37d"></a>

## ipsec.ipsec_tunnel_parameters.site_local_inside_network — ipsec.ipsec_tunnel_parameters.site_local_inside_network / 9c213eabc5cd / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- ipsec.ipsec_tunnel_parameters.site_local_inside_network

<a id="canonical-6bd62f47ea4df47c47596ff98b5825fee0b0ff2f4d6ab03e4ebeab834ff6f2f1"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-dd5debed9630dcd205bcb42d5167f038be3733475774e2503cba24a7597b63be"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.site_local_inside_network / 9c213eabc5cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-202f17e91f8dcbfcfa7ff8300da83b7bdcb040a1b4f0c3fd469b51a02b998e13"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.site_local_inside_network / 9c213eabc5cd / 4

- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-04b993e915a8dad63fef9860ae39edfb47583a7e7de88b234c83682e02e1bd8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5829433fc8dcee8f82928001472fddb6426e332ca1f59ba38ca0833269a6a27"></a>

## ipsec.ipsec_tunnel_parameters.site_local_network — ipsec.ipsec_tunnel_parameters.site_local_network / 906689ae1af6 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- ipsec.ipsec_tunnel_parameters.site_local_network

<a id="canonical-7e443230c5c59923030e6e90c8b786e71c3366556d9bc8e9036b417eccc67f80"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-36cd2cc691b4c67258ed20621ca3f25b7a2faaa4868fd926e9dfcbbeddfafb7b"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.site_local_network / 906689ae1af6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55256dd0e73c535e22786501e22c0450fa9992344d86507991c6ccf62d1df723"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.site_local_network / 906689ae1af6 / 4

- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-63e1300761efb020dbe6ee0af48e1537fc77f867e4aabf6a8b3eaa34f31b4418"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1a8d79823d2753884c9592236549207e8e7f04fda9c79985d73c9a96ab63d28"></a>

## ipsec.ipsec_tunnel_parameters.tunnel_eps — ipsec.ipsec_tunnel_parameters.tunnel_eps / 11603fc5e190 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-15f401740ba27dd7d1b69f64cb5010bd53afa4b2c83be178851d8165b89d60ef)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- ipsec.ipsec_tunnel_parameters.tunnel_eps

<a id="canonical-0c8eb65d0523bf9d6fb4a794857a9381462f5baf3ab2d6b472a00ecb14df1604"></a>

Type: `"list"`. Computed.

Configure tunnel parameters, local and remote IP addresses.

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

<a id="canonical-89279e96cf83c6073a8658b663c2b0e89728785905cd2b7f7d7dd6c7b8169c69"></a>

## Direct properties — ipsec.ipsec_tunnel_parameters.tunnel_eps / 11603fc5e190 / 3

<a id="canonical-d97e6a134898c95f01c9ae63c4980fc2cb2e5409cff14fbbd0d2800cca9dbec2"></a>

<a id="canonical-1c8947750377a7d4367f20989e0432c3bb4506482ff1d01947ded2e5456cdac7"></a>

## interface property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 11603fc5e190 / 4

Type: `"string"`. Computed.

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

<a id="canonical-63dedc161ffc6d347993f5d1e8b389e1eda821bcb33f97c03c0aef0aee7a89e1"></a>

<a id="canonical-50f4c3fe81a8c6536cd89dccef4ef5efaa121190f03f5ff6304dc3eb9ae7b49e"></a>

## local_tunnel_ip property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 11603fc5e190 / 5

Type: `"string"`. Computed.

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

<a id="canonical-cbe4b937fc8a3b5c4c323498ce08d448bcc33db7f09a52f13a5a0f48372335b3"></a>

<a id="canonical-64ad96949ce4743f3d0a61145d83252f84bb282e613382e5fe6f7e4f44053054"></a>

## node property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 11603fc5e190 / 6

Type: `"string"`. Computed.

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

<a id="canonical-c86add8a3d1cec097f7412dcebd5cc318051dfb8b854c505eefb82fb438ad1b0"></a>

<a id="canonical-e71691017e8a9485a7340606b51bc70304b6fe35f93ede2b79ae8c904d34de98"></a>

## remote_tunnel_ip property — ipsec.ipsec_tunnel_parameters.tunnel_eps / 11603fc5e190 / 7

Type: `"string"`. Computed.

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

<a id="canonical-3790f3426495104e0259e95e0a1aa1264048311f8e606eb340cbc12cfc8f6bef"></a>

## Next pages — ipsec.ipsec_tunnel_parameters.tunnel_eps / 11603fc5e190 / 8

- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-63355726d0837996857d62872ba70df7ff1a9c30dc43c252868820690beb3ddf)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
