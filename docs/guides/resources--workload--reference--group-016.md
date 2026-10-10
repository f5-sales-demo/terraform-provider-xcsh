---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3003033200233332-1312033123003330-0333110330231121-1121001133300201-1113031202230011-2202031331030030-0332101313022233-1012103132120022"></a>

## `service.containers.command` property

Type: `["list", "string"]`. Optional.

Command to execute. Overrides the Docker image's ENTRYPOINT.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](resources--workload--reference--group-016.md#canonical-2222010100332200-0330220022103313-0031302313220023-0003301021000001-1132001330133330-0220131323032221-0233002122312201-1330331220112212): complete subsection reference.

- [default_flavor](resources--workload--reference--group-016.md#canonical-3001233032001100-3130331312023200-1301300310120122-1301123133311223-2123211012022020-1331112121220323-2202211100213203-0020000120133312): complete subsection reference.

<a id="canonical-3031113011001222-1021031021023303-0130132000233031-3023023201011203-3123222201323023-2213001310102212-3002132132130333-0331303330110311"></a>

<a id="canonical-2112330223213211-0131333330202203-2333132120022331-0002002000132010-0002311032313300-0130003122102330-1200110130323310-0220302120221302"></a>

## `service.containers.flavor` property

Type: `"string"`. Optional.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Additional upstream details:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](resources--workload--reference--group-016.md#canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212): complete subsection reference.

<a id="canonical-3323033031200113-3030032331023300-1230331021201120-3330031011321033-0131310230333232-1110230121320212-1300021023122121-0013112130303020"></a>

<a id="canonical-0122012123303221-1302202103022101-0003020331010133-2020333210313030-2322020322101211-1001100301222323-0202200023333200-1312013311333300"></a>

## `service.containers.init_container` property

Type: `"bool"`. Optional.

Specialized container that runs before application container and runs to completion.

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

- [liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033): complete subsection reference.

<a id="canonical-1031111210111200-2020323332010033-0223131130102322-2332230303302230-1202322113332021-0330332211033102-2331031011133311-3003231023321111"></a>

<a id="canonical-2220200313001202-2100121110100020-2033322122033231-1013120131312130-3221031103030002-1012310201001100-0113003112213032-3123103020233220"></a>

## `service.containers.name` property

Type: `"string"`. Optional.

Name. Name of the container.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233): complete subsection reference.

<a id="canonical-2222010100332200-0330220022103313-0031302313220023-0003301021000001-1132001330133330-0220131323032221-0233002122312201-1330331220112212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.custom_flavor` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.custom_flavor

<a id="canonical-2330212011000110-2031032322331331-2002033302213221-3332303123223322-0103312032023232-0331103310311201-0010032002023213-1313023022321033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101213301221123-0122120231300130-3221230300021323-2121223100210301-2231133220010320-2133121300300010-3210311322113023-2130320213020003"></a>

### Direct properties for `service.containers.custom_flavor`

<a id="canonical-2210213021332120-3131011112122213-0202031313300311-1023223003030123-0202132202011330-3032330202330313-3021201011100000-1033201023001312"></a>

#### `service.containers.custom_flavor.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2312012222103320-2032201322102300-0301133102211301-3021301213230132-2130123211320310-3002112311330210-2100022331112013-1211303131211303"></a>

<a id="canonical-2213012310123102-3032201212223222-2320133131333103-1031031220211231-1323020011101123-1001303332332132-1203313232000103-3022223303101302"></a>

#### `service.containers.custom_flavor.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3113103200331200-0101222301010111-1030230102232320-2023321020330331-3300313222123130-0222301223122200-2300223123302322-1100230003232201"></a>

<a id="canonical-0123203110222330-2120033332312301-0013113000002213-3321232001120210-3200202111121210-2013020003323002-2102231222102330-2031232230033023"></a>

#### `service.containers.custom_flavor.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3001233032001100-3130331312023200-1301300310120122-1301123133311223-2123211012022020-1331112121220323-2202211100213203-0020000120133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.default_flavor` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.default_flavor

<a id="canonical-3213222203312300-0002322220132313-3001131013201012-1023120001321022-3303010130333100-3211121223203202-2333122302231311-1012013113112021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default flavor.

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
default_flavor = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.image` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.image

<a id="canonical-2310132311110022-1001112000303311-0223131203012103-3023203320332321-0030033013231222-1130303230130033-0121331330202212-2200321121032111"></a>

Type: `"object"`. single nested block, Optional.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322302303000100-2002232102002012-1211011123102200-0323303212122102-3111131121303221-0111100210233022-1233130103121032-1232103121232000"></a>

### Direct properties for `service.containers.image`

- [container_registry](resources--workload--reference--group-016.md#canonical-0021031112302223-0222233313230011-0030112311311132-2103223310130232-3323232123310011-1223212022012321-3223332210202100-2103322001210303): complete subsection reference.

<a id="canonical-1121220301211030-2021132320210201-0000311133120212-2031020320111300-3011333001202201-1131000213323220-2321221201013203-3112133220321201"></a>

<a id="canonical-3112122130231100-0002330100301003-1223222232330223-3311303232320310-2321013230330032-3322201010111123-3033302220212132-3101113321132132"></a>

#### `service.containers.image.name` property

Type: `"string"`. Optional.

Name is a container image which are usually given a name such as alpine, Ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [public](resources--workload--reference--group-016.md#canonical-1001300203132123-0132201313130210-2121111130312013-2201212113112222-1320233221311321-0132232023132312-1111201011323210-0023330021200221): complete subsection reference.

<a id="canonical-3331030011210001-3213221322211022-2302021313131101-3303010203032222-0311011303121120-3013023312013330-2023013301330322-0320002300011330"></a>

<a id="canonical-2021210222200313-3111122103111012-1323103031200332-2212302311132331-3003302001011100-1332112101031313-0312100020202232-3122002210001003"></a>

#### `service.containers.image.pull_policy` property

Type: `"string"`. Optional.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Additional upstream details:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0021031112302223-0222233313230011-0030112311311132-2103223310130232-3323232123310011-1223212022012321-3223332210202100-2103322001210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.image.container_registry` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.image](resources--workload--reference--group-016.md#canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212)
- service.containers.image.container_registry

<a id="canonical-3113001312221102-2213031200033121-2001202220202211-1130330003102310-3113001210003101-1003102032301323-1130212233233101-0323312103330033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000130121213032-0121023311210003-2103223332212230-1120020310221032-2123111210213003-2001213132033303-3113033221033231-3230130221231003"></a>

### Direct properties for `service.containers.image.container_registry`

<a id="canonical-3330313231133200-2313221130323023-2001002333330300-2011303121120121-0102012200121110-0311221201202130-3012222012123220-3312203213311102"></a>

#### `service.containers.image.container_registry.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2101001231323021-3212312011133103-1130213130311320-2221213113032312-1103222102132113-0132013131221332-3021230101203030-1122122332020321"></a>

<a id="canonical-1301000100320230-0033022003303132-2233000323233221-2023121133133302-2331302103303300-1322100202331331-1103031002233323-3012210010203223"></a>

#### `service.containers.image.container_registry.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3222221202323230-1230230123330112-3103020332020123-3002003103213321-2331003210203123-1312201001112311-1230310030100102-1201303103331203"></a>

<a id="canonical-2220220301130223-2233012200310120-1231000203221211-3203301212021002-3220312310332001-1021220000103131-2121200303131311-1130112120032312"></a>

#### `service.containers.image.container_registry.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1001300203132123-0132201313130210-2121111130312013-2201212113112222-1320233221311321-0132232023132312-1111201011323210-0023330021200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.image.public` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.image](resources--workload--reference--group-016.md#canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212)
- service.containers.image.public

<a id="canonical-1022221322321230-0033011020311123-2100230102233033-2132303300200232-0032221213302320-3122312113230203-3300221030132013-1321003222122020"></a>

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
public = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.liveness_check

<a id="canonical-3200312301322013-1001022122222111-1101331103102222-1122320012122131-2012003013310313-1221111000011030-3330211001203133-1301003002231233"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333011232203100-1313232002302110-3333113023121211-3310330010302103-0301303102203200-0103302110303100-0212302122030311-1002232200012103"></a>

### Direct properties for `service.containers.liveness_check`

- [exec_health_check](resources--workload--reference--group-016.md#canonical-3012230211213021-1302303300032103-0333331110210113-0212020331213003-2223223101203330-0131322021233300-0330000033011012-0100031330133333): complete subsection reference.

<a id="canonical-3212212021223210-3011320130130330-3213132121021313-0320021002323132-0333112100300001-0120331012033222-1313213220131211-2231323323332211"></a>

<a id="canonical-2300002313211113-2103031313202033-1322301331213001-0022010211131003-1213312222013113-1022103023022013-2220003033303212-3231331122000223"></a>

#### `service.containers.liveness_check.healthy_threshold` property

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-016.md#canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300): complete subsection reference.

<a id="canonical-2103300022101100-0301200100023333-0133021301303033-0210001313100223-1321131020210021-1113032301031102-3212023130023213-3131330031123033"></a>

<a id="canonical-2233221332013032-1202010101330132-3022012032130010-3032123031123023-3120021223332303-1003323311102001-1313322332323033-2312332231201320"></a>

#### `service.containers.liveness_check.initial_delay` property

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3332213031120213-3033000211201321-3001001020100030-1000200330322122-2230220122023000-2020013033122301-0302220010013301-1011121211030132"></a>

<a id="canonical-3221203223332211-2211032033020103-3031002110131030-3300003120020003-0300303210222222-0312012012202203-3112001022311100-0001002012032301"></a>

#### `service.containers.liveness_check.interval` property

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-016.md#canonical-2103103310121310-1120211222002323-0221023221133310-1322223333330221-2301032221202220-3302011123032012-2321300011332311-2211203122202202): complete subsection reference.

<a id="canonical-3010112221203322-2210030022202211-1012232210011013-2033323202203012-1122003320212021-0001203122321222-2203321303323213-3133120233220101"></a>

<a id="canonical-3301330212033211-1100300212033233-0312032331002112-1212013321312220-2002322121300031-3332313011001310-3122023020300021-3002032202131201"></a>

#### `service.containers.liveness_check.timeout` property

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0001212111230233-0101203201222111-0130200010121010-2111031121322221-0312213010010011-2211302202222300-2112020300102212-1021313230331122"></a>

<a id="canonical-3031211110233112-3002112232023000-2021313333323013-2302233012121032-0022312210231022-3211213312023130-0212111133200000-0300300202110113"></a>

#### `service.containers.liveness_check.unhealthy_threshold` property

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-3012230211213021-1302303300032103-0333331110210113-0212020331213003-2223223101203330-0131322021233300-0330000033011012-0100031330133333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- service.containers.liveness_check.exec_health_check

<a id="canonical-3213001133132100-0022000013331223-0102022212201221-1310123100302100-0202023221121202-0310020011313002-0303231103312222-3030033111133201"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Additional upstream details:

ExecHealthCheckType describes a health check based on "run in container" action.

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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121330323122030-3210321331011122-3301231303201013-2020313001212231-2200133012201300-2330323131301022-2200320111133100-1220302032121123"></a>

### Direct properties for `service.containers.liveness_check.exec_health_check`

<a id="canonical-1330303223121313-0113301220323011-0123023321231211-1130110222120310-3212112332120133-3011221111021010-0200033021010133-0022212122212313"></a>

#### `service.containers.liveness_check.exec_health_check.command` property

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- service.containers.liveness_check.http_health_check

<a id="canonical-0221220000310303-0132021123203210-2230313331320000-3112132012133101-2312200320000231-3011233312020013-3023201003131232-3311230103121311"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330333013120013-3101320311231123-2001222211323111-0003301020332010-3200022112221100-1101110203020022-0113213012120120-1221233200030210"></a>

### Direct properties for `service.containers.liveness_check.http_health_check`

<a id="canonical-1123121231122300-0100212003003213-2231102122002101-2030231311300300-0110302123311313-0220333300220022-3033113120002013-2103012231001121"></a>

#### `service.containers.liveness_check.http_health_check.headers` property

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3120000000003221-2332033130310232-3313232303113020-3131223122303212-3101310011011120-3100332210102320-0220101013131210-0203223000102330"></a>

<a id="canonical-0323200112013301-2201100202333020-2022310223230133-1331302312022123-0323322003201111-3202213033323330-1033333212232102-3132031001112320"></a>

#### `service.containers.liveness_check.http_health_check.host_header` property

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-0311131223223233-0320001330212322-1111121331111032-2001101220200230-1323330312212200-0023321321012321-3332021021303123-0333311303332232"></a>

<a id="canonical-1023323300302122-2321000001302223-3132232001213323-2203133022202031-3122000100120232-1013012302222221-1301322133002303-3013303302013230"></a>

#### `service.containers.liveness_check.http_health_check.path` property

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-016.md#canonical-0032301211110110-3120122131122110-1322202003113123-0133301222022220-1212011232313323-2201303022200113-0330122110130203-1000032333223023): complete subsection reference.

<a id="canonical-0032301211110110-3120122131122110-1322202003113123-0133301222022220-1212011232313323-2201303022200113-0330122110130203-1000032333223023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- [service.containers.liveness_check.http_health_check](resources--workload--reference--group-016.md#canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300)
- service.containers.liveness_check.http_health_check.port

<a id="canonical-0113212303003000-0231302213223003-1223120301320121-2030112022020310-0010103100022122-1130320210030231-3231002011323223-2100331200003012"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230112023112020-2200101211333220-0201122223320022-3201322032321222-1223222032313002-1003022011323330-2121322101002032-0202020211112012"></a>

### Direct properties for `service.containers.liveness_check.http_health_check.port`

<a id="canonical-0022211221333203-3031020200030103-0221013313013111-3023103010222323-1123221312112202-2011331232102001-2330012032012130-1003332012320220"></a>

#### `service.containers.liveness_check.http_health_check.port.name` property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-3321033232031211-3303021130031101-2221231100020332-0232211010200302-3231022221211002-3212111123332120-0313311303222303-3322032020001012"></a>

<a id="canonical-2120132030031222-2221223120032322-1200301100133220-3030100120032031-1301030010230333-1210323100332012-3303202111313203-3231232223210102"></a>

#### `service.containers.liveness_check.http_health_check.port.num` property

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
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

<a id="canonical-2103103310121310-1120211222002323-0221023221133310-1322223333330221-2301032221202220-3302011123032012-2321300011332311-2211203122202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- service.containers.liveness_check.tcp_health_check

<a id="canonical-0022222010210210-0023023223202112-1012203103133131-2233203030232320-2132033130132322-0312312112310131-2020111312111003-3221011230201002"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103200301331230-0002120032103110-2231012010122322-3003333102031311-3311230202322213-2113212301133033-1300110030133121-0300200110020020"></a>

### Direct properties for `service.containers.liveness_check.tcp_health_check`

- [port](resources--workload--reference--group-016.md#canonical-0033013130231022-2310211300202230-3301310301233111-0120011223101333-0032303121201113-3333330023221310-1033121202312312-2323310020123331): complete subsection reference.

<a id="canonical-0033013130231022-2310211300202230-3301310301233111-0120011223101333-0032303121201113-3333330023221310-1033121202312312-2323310020123331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.liveness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- [service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2103103310121310-1120211222002323-0221023221133310-1322223333330221-2301032221202220-3302011123032012-2321300011332311-2211203122202202)
- service.containers.liveness_check.tcp_health_check.port

<a id="canonical-2312120132301322-2013103000103213-1000013000100312-1320332203323130-0331102201013202-0001321311200121-3331123223000023-2331313331021313"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231201013310131-2101310332120023-3003220312322230-3122203232113122-2210333002201131-3320031032202031-2322221021221120-1303003220132323"></a>

### Direct properties for `service.containers.liveness_check.tcp_health_check.port`

<a id="canonical-1333310111223221-3030330113211100-2313110332122231-1220113002232332-0000212200011200-3312321313302021-2032301033111133-1302221100131203"></a>

#### `service.containers.liveness_check.tcp_health_check.port.name` property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-2100010123321120-3113310321123100-3233220212003132-0002103010303321-2201221003330322-2331231121312002-2010300122023021-0313112000320122"></a>

<a id="canonical-2211100111201221-1111032221220332-1333321323002001-2222322023201333-0310112110312022-3313311333121212-2100122023120200-3333120000002210"></a>

#### `service.containers.liveness_check.tcp_health_check.port.num` property

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
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

<a id="canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.readiness_check

<a id="canonical-3002222322321212-1302332323330111-0221323203013230-3232211222032232-1010023230212122-2220220210013221-0112213303301130-1032030033211012"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330212230211231-0113233013312002-0013302132102032-3301201200122002-2200210013320313-2211312332200232-3233012102231223-2110010122102120"></a>

### Direct properties for `service.containers.readiness_check`

- [exec_health_check](resources--workload--reference--group-016.md#canonical-2121212020120232-2033301021030021-0103132112101003-3221002101233112-2232113331320200-1122220030311221-0213111103312110-0220201122030101): complete subsection reference.

<a id="canonical-3303110220333201-2313312220311213-2101111100203010-1113331032311223-2311130001123001-3133310212333021-0300103221223132-1130021313030233"></a>

<a id="canonical-3220303022203123-0103012103102133-0110110302000200-1112323023132222-0211101311132231-2021211110111102-1232220302110120-1332000312111311"></a>

#### `service.containers.readiness_check.healthy_threshold` property

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-016.md#canonical-2002231301213013-0300131302213232-3012321112213220-0202230303103030-3332222333020013-3300112303113211-3301002230313323-1131103111202112): complete subsection reference.

<a id="canonical-1223201321033000-0013012200212102-3202203212202213-3233312331223123-2210003121110300-3210020110001201-1130021213232231-2033122202112130"></a>

<a id="canonical-3010230011011023-3211221012013013-2330112003332330-3231010033332223-1321312103120300-3121300121101200-2230201303013010-2203232121332001"></a>

#### `service.containers.readiness_check.initial_delay` property

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0122000123113002-1222123200123222-0203113031312301-1031020301311322-3000123111003302-0133120023212033-3222130013131123-2333231300111122"></a>

<a id="canonical-0000031030030331-3201320110030301-0223213031323300-3112130011313231-3323213022213123-1020112023311123-1032301311211331-3123020111202032"></a>

#### `service.containers.readiness_check.interval` property

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-016.md#canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323): complete subsection reference.

<a id="canonical-2320020210331301-2133001321012023-1112113232232212-1022110002023320-1220313311010232-3123030322223111-0332102131100223-0130230102200131"></a>

<a id="canonical-0032013222000133-2220303332121222-0201311311003103-3200203100233023-2220231310320233-0303302011213130-1213210312121023-1022213001221112"></a>

#### `service.containers.readiness_check.timeout` property

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1031310101132333-0031311322003221-3122031132223200-2303030231020001-3320213102111301-0013012303133113-2211001332121210-1302301123333200"></a>

<a id="canonical-0013022213112311-3103330313022211-3113022030003223-3301301011221232-2103331311200211-1120102030021201-0223303330213302-2130012032110212"></a>

#### `service.containers.readiness_check.unhealthy_threshold` property

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2121212020120232-2033301021030021-0103132112101003-3221002101233112-2232113331320200-1122220030311221-0213111103312110-0220201122030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- service.containers.readiness_check.exec_health_check

<a id="canonical-1321101121232121-1103312120310203-2212301210110112-3110030233322322-2100131230313120-1031033120321113-3313313301221101-0030001300222031"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Additional upstream details:

ExecHealthCheckType describes a health check based on "run in container" action.

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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100303202320331-3001011300112331-0300322330013020-0312021330012332-1220213320112133-0303100112130303-2112331113202111-0203023200130110"></a>

### Direct properties for `service.containers.readiness_check.exec_health_check`

<a id="canonical-2132313010310330-2020120310301200-2122122301010011-1132302331311232-0221300212212020-2103103111301012-0100301210101320-2320022131001203"></a>

#### `service.containers.readiness_check.exec_health_check.command` property

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2002231301213013-0300131302213232-3012321112213220-0202230303103030-3332222333020013-3300112303113211-3301002230313323-1131103111202112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- service.containers.readiness_check.http_health_check

<a id="canonical-3231000012130123-0033320101300313-2222103310303120-1012300031333110-3300130320223132-3233122232101233-2313013330033321-1000321011302311"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130013203111131-3123221202102213-1303022123111213-3222002323332031-2013213110121112-3322232010300312-3123231023230123-1212203332321111"></a>

### Direct properties for `service.containers.readiness_check.http_health_check`

<a id="canonical-3122021310313112-2213132302200133-1211213220313100-0211213300203300-2303311121330122-3110222222202203-2303303121223320-3002021003031011"></a>

#### `service.containers.readiness_check.http_health_check.headers` property

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0212012033021300-2332313032010223-0102022230102203-0311231322033302-0222203121320101-3122020202113120-3031200330020133-0323030131003331"></a>

<a id="canonical-2233000202101332-3131220332012222-2231003101202133-3012333113311131-2312212310213030-0121002003331232-1330312321203102-3102332110032103"></a>

#### `service.containers.readiness_check.http_health_check.host_header` property

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-1122322022333111-3032030332302302-3331212031320311-3023031311023313-0202320300211122-1033020132223223-2330300031100320-0212230201130110"></a>

<a id="canonical-1010020203222222-2100220311321132-0203133202321200-1300003301030323-3103313323132220-1002201023222023-1033220301313002-2120033312002223"></a>

#### `service.containers.readiness_check.http_health_check.path` property

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-016.md#canonical-0210321022333123-2321122000201010-2122130303010002-1212010310030122-3332222320112012-0100000032032322-1231111010210313-2311132203031320): complete subsection reference.

<a id="canonical-0210321022333123-2321122000201010-2122130303010002-1212010310030122-3332222320112012-0100000032032322-1231111010210313-2311132203031320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- [service.containers.readiness_check.http_health_check](resources--workload--reference--group-016.md#canonical-2002231301213013-0300131302213232-3012321112213220-0202230303103030-3332222333020013-3300112303113211-3301002230313323-1131103111202112)
- service.containers.readiness_check.http_health_check.port

<a id="canonical-3023031203133323-1323220032103222-3300213102323020-1321210233103302-1311333201312310-2233111022222322-3010320110120203-2303023023201322"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213233321001132-2313020130102300-0031022103300322-0200202310020221-2303212103130021-3211023310130320-2200202100100102-2102030310212333"></a>

### Direct properties for `service.containers.readiness_check.http_health_check.port`

<a id="canonical-3311113310000130-2013202023122010-2330101102302233-3002031320022233-3303322033020330-1212303212111111-3310211333131223-0331231030213011"></a>

#### `service.containers.readiness_check.http_health_check.port.name` property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-1022103023032121-0013013110310233-3103221330110013-2303123013310133-1022210310301033-1010230030311120-1212011123221211-2013333212023120"></a>

<a id="canonical-1121213132310031-0211110012010132-0332033020233121-0331131300200112-3201201033101002-3002131121033103-1022210000012130-2011212130031001"></a>

#### `service.containers.readiness_check.http_health_check.port.num` property

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
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

<a id="canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- service.containers.readiness_check.tcp_health_check

<a id="canonical-3212313330021302-0320212212321000-3020322033302203-0330022102102000-3311210120302000-3211320201201003-0321001112122113-2222122201331003"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230023133232220-0003303120110303-0303132310020203-1323112333302022-3201001310131202-0221220221011102-2323221200000011-0311100132322312"></a>

### Direct properties for `service.containers.readiness_check.tcp_health_check`

- [port](resources--workload--reference--group-016.md#canonical-2212311102201300-1321133023202333-1202331302220201-3300113313012031-3212022302113322-0323201201231011-1331022002230220-1131231303332213): complete subsection reference.

<a id="canonical-2212311102201300-1321133023202333-1202331302220201-3300113313012031-3212022302113322-0323201201231011-1331022002230220-1131231303332213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.containers.readiness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- [service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323)
- service.containers.readiness_check.tcp_health_check.port

<a id="canonical-0123122132011210-0313023320103313-2112302113011303-3230023212120223-2023003312200112-1101133023312213-3332323032100232-2232311020001320"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333233210320003-2301100100233003-0221210111032100-0203232132233022-1022232103131331-2002010312203223-0232113212100322-3011020222103203"></a>

### Direct properties for `service.containers.readiness_check.tcp_health_check.port`

<a id="canonical-1213333210320330-0121211221123110-0301312311101012-1322320232322330-2012100232233312-3013231123302002-0233002001311221-3123130223203131"></a>

#### `service.containers.readiness_check.tcp_health_check.port.name` property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-2131120120031030-2121122230012032-0010013232301303-1132332012130100-1031303012001003-0210133332001003-0100320321233100-3000222213031101"></a>

<a id="canonical-2210231310032112-0011330232233220-3000200013232133-2231301232000110-3002311002211233-2303230011131221-1231113330021222-1311131333032231"></a>

#### `service.containers.readiness_check.tcp_health_check.port.num` property

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
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

<a id="canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.deploy_options

<a id="canonical-2003022010122311-3321021022020131-2103000013201023-1222302021013313-3001303103030311-1332132113321232-2300111213130322-3321232233101331"></a>

Type: `"object"`. single nested block, Optional.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

Terraform syntax:

```terraform
deploy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032010020312200-2121131321210330-2323110003203001-2023303033223232-0120033020132022-2331023002032132-1001211311110212-1301312100113233"></a>

### Direct properties for `service.deploy_options`

- [all_res](resources--workload--reference--group-016.md#canonical-3320122331212021-2103131122203112-1021030123303133-1113130130300120-1332001011321111-1021203031103220-1133320202311232-2230100023300313): complete subsection reference.

- [default_virtual_sites](resources--workload--reference--group-016.md#canonical-0123021220301230-0122112311122001-3213103300231202-3110232110313113-2030320231100102-1111033100101031-3212301103121201-1001132210103332): complete subsection reference.

- [deploy_ce_sites](resources--workload--reference--group-016.md#canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213): complete subsection reference.

- [deploy_re_sites](resources--workload--reference--group-016.md#canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--reference--group-017.md#canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202): complete subsection reference.

<a id="canonical-3320122331212021-2103131122203112-1021030123303133-1113130130300120-1332001011321111-1021203031103220-1133320202311232-2230100023300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.all_res` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.all_res

<a id="canonical-1321313220330303-2030200020321220-2201323120232121-3001200001302102-2212210330130213-3210010332323111-0030302222202312-2323123020220300"></a>

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
all_res = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123021220301230-0122112311122001-3213103300231202-3110232110313113-2030320231100102-1111033100101031-3212301103121201-1001132210103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.default_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.default_virtual_sites

<a id="canonical-1321212010121323-3300132300321312-1223311311002222-1012312001300000-0032330033120033-0003032322003211-3301133200012103-3003112022230201"></a>

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
default_virtual_sites = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_ce_sites

<a id="canonical-2121110312021123-3303032202132020-0101132103323032-2213131023102202-2332222102323102-3002131213322233-0223101120331133-2311311023101112"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to deploy a workload on specific Customer sites.

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
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100132213233323-0130321121321113-1131332022302000-3020132130222003-0322301130021303-3331121320012332-0102010200320123-3332032132121011"></a>

### Direct properties for `service.deploy_options.deploy_ce_sites`

- [site](resources--workload--reference--group-016.md#canonical-3101000002123000-1233232311011210-3101322210100023-3232310031032321-1123000020111333-3201031203033230-1312221132032301-1322320221213110): complete subsection reference.

<a id="canonical-3101000002123000-1233232311011210-3101322210100023-3232310031032321-1123000020111333-3201031203033230-1312221132032301-1322320221213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_ce_sites](resources--workload--reference--group-016.md#canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323)
- service.deploy_options.deploy_ce_sites.site

<a id="canonical-3022112202330230-2323132320333020-3310232310031222-1033022200310011-3011313033220232-2001311210030323-2320203202110223-0313323113111323"></a>

Type: `"object"`. list nested block, Optional.

Which customer sites should this workload be deployed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313311032300130-3332302230110300-3130310230120021-1333212133010221-1330132312300320-0121201301222223-2200210101230000-1101120200020110"></a>

### Direct properties for `service.deploy_options.deploy_ce_sites.site`

<a id="canonical-1300310032022010-2301202212203020-2001231010011123-1130311313111123-3023122100133113-0202033232311011-2320322130222230-3121111022220122"></a>

#### `service.deploy_options.deploy_ce_sites.site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3003221210321002-3132210000331333-0301120321222221-0012322302221122-3101323001000121-3012003303221313-2100130021032213-2302130011320031"></a>

<a id="canonical-3201033302301110-0211110310031222-3312111013111122-0021311033211220-1133211132323000-1131131310313200-1022203223121003-2232010113321123"></a>

#### `service.deploy_options.deploy_ce_sites.site.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3011221310022020-3313233302330111-0320032032330032-2201133333331013-1300310231131202-0321230313021303-3022313222103130-2223112210221323"></a>

<a id="canonical-2013223021003322-1011002220210133-2210000102212112-2111223132311331-3122303122032213-0311033233222101-1222223103301312-3213211102223012"></a>

#### `service.deploy_options.deploy_ce_sites.site.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-3220020110313002-1101331223200023-0333220300121021-0302330110232100-2122233122230312-1312201231103021-2201020012022111-1030100333101110"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to deploy a workload on specific Customer virtual sites.

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
deploy_ce_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130032131113101-0332330101230130-0003133223002021-2123221330201321-1031221232113002-3213302221211332-2320132231232200-1222011012313220"></a>

### Direct properties for `service.deploy_options.deploy_ce_virtual_sites`

- [virtual_site](resources--workload--reference--group-016.md#canonical-1201201122302303-1001333103123313-0130101233033100-3200113022111130-0213333133332302-3100202312211131-2120213320312021-0332232100003110): complete subsection reference.

<a id="canonical-1201201122302303-1001333103123313-0130101233033100-3200113022111130-0213333133332302-3100202312211131-2120213320312021-0332232100003110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_ce_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213)
- service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-2132123023233221-2023233213131021-0201022213110220-0132110011200013-1223023220112011-2013133033221121-1002231221113201-2110232001223123"></a>

Type: `"object"`. list nested block, Optional.

Which customer virtual sites should this workload be deployed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203130212000130-2120203132100331-1210200200332321-0303032021232212-1012203121332001-1211203213330303-3030310000002322-3103312221022202"></a>

### Direct properties for `service.deploy_options.deploy_ce_virtual_sites.virtual_site`

<a id="canonical-2223103130011133-1210222212312232-0321320102102323-1003332100131223-1122230210300013-2301212121100112-3103111302102121-0132122121002222"></a>

#### `service.deploy_options.deploy_ce_virtual_sites.virtual_site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0000201200101033-3311031231220200-0120211203211321-3300131021312102-0201312300201321-2303333033013303-1333203021032001-0321001210021002"></a>

<a id="canonical-1322212312022122-2222313011202200-1313231303332320-1232213200211213-1330332320322120-1312102100102231-1222022221302130-3103300202113203"></a>

#### `service.deploy_options.deploy_ce_virtual_sites.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2232202030332313-1022232010113303-3301030202122221-0211320103323103-2101201130030030-3231331220013202-2323221011320303-2101212100011212"></a>

<a id="canonical-2221333022101020-2122330101123230-2031301123003311-2301211100323222-0122233312023202-3231201310112203-2133103330201223-1233033022011122"></a>

#### `service.deploy_options.deploy_ce_virtual_sites.virtual_site.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.deploy_options.deploy_re_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_re_sites

<a id="canonical-3013322233213022-2310103223230200-1303031300003322-3323012231322013-0211003110301321-2220020221322210-1330311000100132-3332112313001333"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to deploy a workload on specific Regional Edge sites.

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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003333121220210-2312100313323110-1310132321033301-1320220000133032-3322122211300210-0130002130121003-2222100002131201-3201331010100313"></a>

### Direct properties for `service.deploy_options.deploy_re_sites`

- [site](resources--workload--reference--group-017.md#canonical-2033131030312023-2200112222221310-2222232223012011-0200121201203231-3033103333011032-0331022200133101-2230032301132121-2012311230312312): complete subsection reference.
