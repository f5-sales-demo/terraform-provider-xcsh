---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2101213301221123-0122120231300130-3221230300021323-2121223100210301-2231133220010320-2133121300300010-3210311322113023-2130320213020003"></a>

## service.containers.custom_flavor — custom_flavor / 330212223311 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.custom_flavor

<a id="canonical-2330212011000110-2031032322331331-2002033302213221-3332303123223322-0103312032023232-0331103310311201-0010032002023213-1313023022321033"></a>

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
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213012310123102-3032201212223222-2320133131333103-1031031220211231-1323020011101123-1001303332332132-1203313232000103-3022223303101302"></a>

## Direct properties — custom_flavor / 330212223311 / 3

<a id="canonical-2210213021332120-3131011112122213-0202031313300311-1023223003030123-0202132202011330-3032330202330313-3021201011100000-1033201023001312"></a>

<a id="canonical-0123203110222330-2120033332312301-0013113000002213-3321232001120210-3200202111121210-2013020003323002-2102231222102330-2031232230033023"></a>

## name property — custom_flavor / 330212223311 / 4

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

<a id="canonical-2312012222103320-2032201322102300-0301133102211301-3021301213230132-2130123211320310-3002112311330210-2100022331112013-1211303131211303"></a>

<a id="canonical-2203011331033022-0333231130023232-0202023330131131-0222311322212300-1022033320030210-2223320133232030-2013100010203201-3303011020130210"></a>

## namespace property — custom_flavor / 330212223311 / 5

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

<a id="canonical-3113103200331200-0101222301010111-1030230102232320-2023321020330331-3300313222123130-0222301223122200-2300223123302322-1100230003232201"></a>

<a id="canonical-2002002101332103-0002110201011021-2220210322112201-3202200113030030-0213200003003233-2303333323012002-2100131311010001-3012330020130113"></a>

## tenant property — custom_flavor / 330212223311 / 6

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

<a id="canonical-1002022112210023-0321333331232021-2333003210122013-1230031112213122-1313212113111223-2213313202233111-3200020330031100-1323221120233201"></a>

## Next pages — custom_flavor / 330212223311 / 7

- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3001233032001100-3130331312023200-1301300310120122-1301123133311223-2123211012022020-1331112121220323-2202211100213203-0020000120133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032120223203100-1121312301312333-3122303000110123-0211200113232021-2031003312120002-0010203022123101-0002011011330200-0223023200220330"></a>

## service.containers.default_flavor — default_flavor / 001312131110 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.default_flavor

<a id="canonical-3213222203312300-0002322220132313-3001131013201012-1023120001321022-3303010130333100-3211121223203202-2333122302231311-1012013113112021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default flavor.

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
default_flavor = {}
```

<a id="canonical-0101221212231202-2031011322212033-2101002232120201-0021033120113200-3021210113203321-2131121030033032-0123330200200033-2003331123101321"></a>

## Direct properties — default_flavor / 001312131110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030000013320010-2032032121131321-1220203122230213-1010122132211010-0233330302322002-2223302330321301-3311000221002201-3323003312221110"></a>

## Next pages — default_flavor / 001312131110 / 4

- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322302303000100-2002232102002012-1211011123102200-0323303212122102-3111131121303221-0111100210233022-1233130103121032-1232103121232000"></a>

## service.containers.image — image / 122122301030 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.image

<a id="canonical-2310132311110022-1001112000303311-0223131203012103-3023203320332321-0030033013231222-1130303230130033-0121331330202212-2200321121032111"></a>

Type: `"object"`. single nested block, Optional.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("container_registry",
    "public")}
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
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112122130231100-0002330100301003-1223222232330223-3311303232320310-2321013230330032-3322201010111123-3033302220212132-3101113321132132"></a>

## Direct properties — image / 122122301030 / 3

- [container_registry](resources--workload--reference--group-016.md#canonical-0021031112302223-0222233313230011-0030112311311132-2103223310130232-3323232123310011-1223212022012321-3223332210202100-2103322001210303): complete subsection reference.

<a id="canonical-1121220301211030-2021132320210201-0000311133120212-2031020320111300-3011333001202201-1131000213323220-2321221201013203-3112133220321201"></a>

<a id="canonical-2021210222200313-3111122103111012-1323103031200332-2212302311132331-3003302001011100-1332112101031313-0312100020202232-3122002210001003"></a>

## name property — image / 122122301030 / 4

Type: `"string"`. Optional.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1020311323201020-0213310002002321-2100100332131300-0202003010313023-1013032332231232-3332112022133212-0202000200332232-3302012221122100"></a>

## pull_policy property — image / 122122301030 / 5

Type: `"string"`. Optional.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"),
}
```

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

<a id="canonical-0322320331332221-1031112222031210-2313002103332023-1333230013200102-1010222303120120-1130303132000001-0122103023103100-0222320220013220"></a>

## Next pages — image / 122122301030 / 6

- [service.containers.image.container_registry](resources--workload--reference--group-016.md#canonical-0021031112302223-0222233313230011-0030112311311132-2103223310130232-3323232123310011-1223212022012321-3223332210202100-2103322001210303)
- [service.containers.image.public](resources--workload--reference--group-016.md#canonical-1001300203132123-0132201313130210-2121111130312013-2201212113112222-1320233221311321-0132232023132312-1111201011323210-0023330021200221)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0021031112302223-0222233313230011-0030112311311132-2103223310130232-3323232123310011-1223212022012321-3223332210202100-2103322001210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000130121213032-0121023311210003-2103223332212230-1120020310221032-2123111210213003-2001213132033303-3113033221033231-3230130221231003"></a>

## service.containers.image.container_registry — container_registry / 230320020221 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.image](resources--workload--reference--group-016.md#canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212)
- service.containers.image.container_registry

<a id="canonical-3113001312221102-2213031200033121-2001202220202211-1130330003102310-3113001210003101-1003102032301323-1130212233233101-0323312103330033"></a>

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
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301000100320230-0033022003303132-2233000323233221-2023121133133302-2331302103303300-1322100202331331-1103031002233323-3012210010203223"></a>

## Direct properties — container_registry / 230320020221 / 3

<a id="canonical-3330313231133200-2313221130323023-2001002333330300-2011303121120121-0102012200121110-0311221201202130-3012222012123220-3312203213311102"></a>

<a id="canonical-2220220301130223-2233012200310120-1231000203221211-3203301212021002-3220312310332001-1021220000103131-2121200303131311-1130112120032312"></a>

## name property — container_registry / 230320020221 / 4

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

<a id="canonical-2101001231323021-3212312011133103-1130213130311320-2221213113032312-1103222102132113-0132013131221332-3021230101203030-1122122332020321"></a>

<a id="canonical-2233200321320132-2220231003212323-3230000020223103-2213012312133220-1101312322222122-1000113132020210-1111021233023321-1332330020033211"></a>

## namespace property — container_registry / 230320020221 / 5

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

<a id="canonical-3222221202323230-1230230123330112-3103020332020123-3002003103213321-2331003210203123-1312201001112311-1230310030100102-1201303103331203"></a>

<a id="canonical-3123322233303030-3020020002031030-3223202333133322-1300230132311231-2300221101131122-2333130320123330-1123303120101032-0133023003001222"></a>

## tenant property — container_registry / 230320020221 / 6

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

<a id="canonical-0032103310331320-3123300302312323-3223131020300210-2301013101212030-3332200203310322-1302202111202322-0221131032200121-0202112102313302"></a>

## Next pages — container_registry / 230320020221 / 7

- [service.containers.image](resources--workload--reference--group-016.md#canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1001300203132123-0132201313130210-2121111130312013-2201212113112222-1320233221311321-0132232023132312-1111201011323210-0023330021200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132202313110310-2123003312231102-1101301311321323-2322223032031202-3031133333012331-2001112210002222-0220113132313332-3302113321012301"></a>

## service.containers.image.public — public / 213130131232 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.image](resources--workload--reference--group-016.md#canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212)
- service.containers.image.public

<a id="canonical-1022221322321230-0033011020311123-2100230102233033-2132303300200232-0032221213302320-3122312113230203-3300221030132013-1321003222122020"></a>

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
public = {}
```

<a id="canonical-2302101213220300-1123322212210310-2221000022312201-1310102003011101-1133311103102323-2201230212313311-1221010132101132-3232211031112211"></a>

## Direct properties — public / 213130131232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211030221000022-2001031210122312-1123132321023023-3021310111111132-1013002001303023-2001233232022100-0133103122312213-0133322021202203"></a>

## Next pages — public / 213130131232 / 4

- [service.containers.image](resources--workload--reference--group-016.md#canonical-1111210320310031-3203221222232112-2101301202201220-1132023000321231-3301101223203031-3300202121122031-1022201320110031-3032000331210212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333011232203100-1313232002302110-3333113023121211-3310330010302103-0301303102203200-0103302110303100-0212302122030311-1002232200012103"></a>

## service.containers.liveness_check — liveness_check / 323321232330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.liveness_check

<a id="canonical-3200312301322013-1001022122222111-1101331103102222-1122320012122131-2012003013310313-1221111000011030-3330211001203133-1301003002231233"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300002313211113-2103031313202033-1322301331213001-0022010211131003-1213312222013113-1022103023022013-2220003033303212-3231331122000223"></a>

## Direct properties — liveness_check / 323321232330 / 3

- [exec_health_check](resources--workload--reference--group-016.md#canonical-3012230211213021-1302303300032103-0333331110210113-0212020331213003-2223223101203330-0131322021233300-0330000033011012-0100031330133333): complete subsection reference.

<a id="canonical-3212212021223210-3011320130130330-3213132121021313-0320021002323132-0333112100300001-0120331012033222-1313213220131211-2231323323332211"></a>

<a id="canonical-2233221332013032-1202010101330132-3022012032130010-3032123031123023-3120021223332303-1003323311102001-1313322332323033-2312332231201320"></a>

## healthy_threshold property — liveness_check / 323321232330 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3221203223332211-2211032033020103-3031002110131030-3300003120020003-0300303210222222-0312012012202203-3112001022311100-0001002012032301"></a>

## initial_delay property — liveness_check / 323321232330 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3332213031120213-3033000211201321-3001001020100030-1000200330322122-2230220122023000-2020013033122301-0302220010013301-1011121211030132"></a>

<a id="canonical-3301330212033211-1100300212033233-0312032331002112-1212013321312220-2002322121300031-3332313011001310-3122023020300021-3002032202131201"></a>

## interval property — liveness_check / 323321232330 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-3031211110233112-3002112232023000-2021313333323013-2302233012121032-0022312210231022-3211213312023130-0212111133200000-0300300202110113"></a>

## timeout property — liveness_check / 323321232330 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-3330203012212313-3233122222033231-0332330021120132-1213131010100211-0332221203020110-2000201003203300-0211000132112313-0103332101000013"></a>

## unhealthy_threshold property — liveness_check / 323321232330 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1313112222313230-0332132021100203-1302213020130113-1131031311103220-1320111013232011-2133312233020133-0230332313230030-0001302321322233"></a>

## Next pages — liveness_check / 323321232330 / 9

- [service.containers.liveness_check.exec_health_check](resources--workload--reference--group-016.md#canonical-3012230211213021-1302303300032103-0333331110210113-0212020331213003-2223223101203330-0131322021233300-0330000033011012-0100031330133333)
- [service.containers.liveness_check.http_health_check](resources--workload--reference--group-016.md#canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300)
- [service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2103103310121310-1120211222002323-0221023221133310-1322223333330221-2301032221202220-3302011123032012-2321300011332311-2211203122202202)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3012230211213021-1302303300032103-0333331110210113-0212020331213003-2223223101203330-0131322021233300-0330000033011012-0100031330133333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121330323122030-3210321331011122-3301231303201013-2020313001212231-2200133012201300-2330323131301022-2200320111133100-1220302032121123"></a>

## service.containers.liveness_check.exec_health_check — exec_health_check / 001333103133 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- service.containers.liveness_check.exec_health_check

<a id="canonical-3213001133132100-0022000013331223-0102022212201221-1310123100302100-0202023221121202-0310020011313002-0303231103312222-3030033111133201"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323332020311100-3210202111132010-3013022033010021-2020010333201112-2331003010332110-1101310103132013-3202101233102121-1000311221202012"></a>

## Direct properties — exec_health_check / 001333103133 / 3

<a id="canonical-1330303223121313-0113301220323011-0123023321231211-1130110222120310-3212112332120133-3011221111021010-0200033021010133-0022212122212313"></a>

<a id="canonical-1213302131131102-0103231301003213-0012022231221212-1000020200030120-2302032333112021-1230322002322230-1310203201332120-3013212110323332"></a>

## command property — exec_health_check / 001333103133 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2122102212222230-2220101003102300-2300111230203101-3022122313103120-2123302212013322-3103233122030110-3021122000331202-0321021101231133"></a>

## Next pages — exec_health_check / 001333103133 / 5

- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330333013120013-3101320311231123-2001222211323111-0003301020332010-3200022112221100-1101110203020022-0113213012120120-1221233200030210"></a>

## service.containers.liveness_check.http_health_check — http_health_check / 102322301211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- service.containers.liveness_check.http_health_check

<a id="canonical-0221220000310303-0132021123203210-2230313331320000-3112132012133101-2312200320000231-3011233312020013-3023201003131232-3311230103121311"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323200112013301-2201100202333020-2022310223230133-1331302312022123-0323322003201111-3202213033323330-1033333212232102-3132031001112320"></a>

## Direct properties — http_health_check / 102322301211 / 3

<a id="canonical-1123121231122300-0100212003003213-2231102122002101-2030231311300300-0110302123311313-0220333300220022-3033113120002013-2103012231001121"></a>

<a id="canonical-1023323300302122-2321000001302223-3132232001213323-2203133022202031-3122000100120232-1013012302222221-1301322133002303-3013303302013230"></a>

## headers property — http_health_check / 102322301211 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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

<a id="canonical-0102110011320200-2202210011122002-3101012032120210-0113310122101002-3010320311230130-1220203333111021-3330321220010300-1002321210211233"></a>

## host_header property — http_health_check / 102322301211 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

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

<a id="canonical-1101001031231111-1300332102212032-1002132013000103-2302030102230023-3113100332310313-3121232030001300-3032322223302320-1312022133130123"></a>

## path property — http_health_check / 102322301211 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3210100332312210-0321302100330132-0222003110211323-2132031210000323-3113211031331311-0322020012213332-3110301323300220-1321301102021311"></a>

## Next pages — http_health_check / 102322301211 / 7

- [service.containers.liveness_check.http_health_check.port](resources--workload--reference--group-016.md#canonical-0032301211110110-3120122131122110-1322202003113123-0133301222022220-1212011232313323-2201303022200113-0330122110130203-1000032333223023)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0032301211110110-3120122131122110-1322202003113123-0133301222022220-1212011232313323-2201303022200113-0330122110130203-1000032333223023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230112023112020-2200101211333220-0201122223320022-3201322032321222-1223222032313002-1003022011323330-2121322101002032-0202020211112012"></a>

## service.containers.liveness_check.http_health_check.port — port / 023013312310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- [service.containers.liveness_check.http_health_check](resources--workload--reference--group-016.md#canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300)
- service.containers.liveness_check.http_health_check.port

<a id="canonical-0113212303003000-0231302213223003-1223120301320121-2030112022020310-0010103100022122-1130320210030231-3231002011323223-2100331200003012"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120132030031222-2221223120032322-1200301100133220-3030100120032031-1301030010230333-1210323100332012-3303202111313203-3231232223210102"></a>

## Direct properties — port / 023013312310 / 3

<a id="canonical-0022211221333203-3031020200030103-0221013313013111-3023103010222323-1123221312112202-2011331232102001-2330012032012130-1003332012320220"></a>

<a id="canonical-2131233222112232-2312113030011010-3221212101002332-3132001023331200-3101033212302101-2111013300233103-2123110200103100-1232102121212222"></a>

## name property — port / 023013312310 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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

<a id="canonical-0213013030332203-0010302132001132-0311313112220000-3320033332223313-0101312020030110-0122101332310301-2222101110310311-2303021332111021"></a>

## num property — port / 023013312310 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3020211230320131-3110100032202230-2031222203200332-0122231202020123-1013230310010132-0233100333232001-0020023122201302-0303130223330011"></a>

## Next pages — port / 023013312310 / 6

- [service.containers.liveness_check.http_health_check](resources--workload--reference--group-016.md#canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2103103310121310-1120211222002323-0221023221133310-1322223333330221-2301032221202220-3302011123032012-2321300011332311-2211203122202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103200301331230-0002120032103110-2231012010122322-3003333102031311-3311230202322213-2113212301133033-1300110030133121-0300200110020020"></a>

## service.containers.liveness_check.tcp_health_check — tcp_health_check / 121031233113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
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

<a id="canonical-3101112312012102-2322023301232333-2002113023213321-2013331023210132-2022133010123310-3310112132322130-0031322111011332-0121220011130002"></a>

## Direct properties — tcp_health_check / 121031233113 / 3

- [port](resources--workload--reference--group-016.md#canonical-0033013130231022-2310211300202230-3301310301233111-0120011223101333-0032303121201113-3333330023221310-1033121202312312-2323310020123331): complete subsection reference.

<a id="canonical-2310112310103113-0320001303320032-1313112030230222-3133120113103023-2311123301112230-1202310302333000-2123022101223002-2033131032331203"></a>

## Next pages — tcp_health_check / 121031233113 / 4

- [service.containers.liveness_check.tcp_health_check.port](resources--workload--reference--group-016.md#canonical-0033013130231022-2310211300202230-3301310301233111-0120011223101333-0032303121201113-3333330023221310-1033121202312312-2323310020123331)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0033013130231022-2310211300202230-3301310301233111-0120011223101333-0032303121201113-3333330023221310-1033121202312312-2323310020123331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231201013310131-2101310332120023-3003220312322230-3122203232113122-2210333002201131-3320031032202031-2322221021221120-1303003220132323"></a>

## service.containers.liveness_check.tcp_health_check.port — port / 113121213103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.liveness_check](resources--workload--reference--group-016.md#canonical-2130113101101012-3022213001032320-0233102123032103-1302200230313311-1222131103210233-2232333201302321-0021320231032101-0322330022002033)
- [service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2103103310121310-1120211222002323-0221023221133310-1322223333330221-2301032221202220-3302011123032012-2321300011332311-2211203122202202)
- service.containers.liveness_check.tcp_health_check.port

<a id="canonical-2312120132301322-2013103000103213-1000013000100312-1320332203323130-0331102201013202-0001321311200121-3331123223000023-2331313331021313"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211100111201221-1111032221220332-1333321323002001-2222322023201333-0310112110312022-3313311333121212-2100122023120200-3333120000002210"></a>

## Direct properties — port / 113121213103 / 3

<a id="canonical-1333310111223221-3030330113211100-2313110332122231-1220113002232332-0000212200011200-3312321313302021-2032301033111133-1302221100131203"></a>

<a id="canonical-1102123331020033-2103032201121032-0013132201220021-1312112323132233-1201211223000202-0102303322120122-1300320123310303-3011122210030303"></a>

## name property — port / 113121213103 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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

<a id="canonical-0123133220012230-1120030032130113-1110201003233131-2131303011011131-3303022022302030-0023333313303210-1000331112032003-2103213231232302"></a>

## num property — port / 113121213103 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1321002333311321-2122322123320120-1202101230113121-0013010330301123-0200331223303001-2210310022030130-0003030221013202-3231110130321100"></a>

## Next pages — port / 113121213103 / 6

- [service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2103103310121310-1120211222002323-0221023221133310-1322223333330221-2301032221202220-3302011123032012-2321300011332311-2211203122202202)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330212230211231-0113233013312002-0013302132102032-3301201200122002-2200210013320313-2211312332200232-3233012102231223-2110010122102120"></a>

## service.containers.readiness_check — readiness_check / 110330101223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- service.containers.readiness_check

<a id="canonical-3002222322321212-1302332323330111-0221323203013230-3232211222032232-1010023230212122-2220220210013221-0112213303301130-1032030033211012"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220303022203123-0103012103102133-0110110302000200-1112323023132222-0211101311132231-2021211110111102-1232220302110120-1332000312111311"></a>

## Direct properties — readiness_check / 110330101223 / 3

- [exec_health_check](resources--workload--reference--group-016.md#canonical-2121212020120232-2033301021030021-0103132112101003-3221002101233112-2232113331320200-1122220030311221-0213111103312110-0220201122030101): complete subsection reference.

<a id="canonical-3303110220333201-2313312220311213-2101111100203010-1113331032311223-2311130001123001-3133310212333021-0300103221223132-1130021313030233"></a>

<a id="canonical-3010230011011023-3211221012013013-2330112003332330-3231010033332223-1321312103120300-3121300121101200-2230201303013010-2203232121332001"></a>

## healthy_threshold property — readiness_check / 110330101223 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0000031030030331-3201320110030301-0223213031323300-3112130011313231-3323213022213123-1020112023311123-1032301311211331-3123020111202032"></a>

## initial_delay property — readiness_check / 110330101223 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0122000123113002-1222123200123222-0203113031312301-1031020301311322-3000123111003302-0133120023212033-3222130013131123-2333231300111122"></a>

<a id="canonical-0032013222000133-2220303332121222-0201311311003103-3200203100233023-2220231310320233-0303302011213130-1213210312121023-1022213001221112"></a>

## interval property — readiness_check / 110330101223 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-0013022213112311-3103330313022211-3113022030003223-3301301011221232-2103331311200211-1120102030021201-0223303330213302-2130012032110212"></a>

## timeout property — readiness_check / 110330101223 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-1011101121110011-1213320311303231-3213002222211310-0123003231210010-1110332101023021-1203111331322030-1101303320130123-2333122331121320"></a>

## unhealthy_threshold property — readiness_check / 110330101223 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1103110102210022-2213221131320110-2011010203011110-0300201132301301-3220232313311000-2303132000301130-2331031332211202-2003000211002213"></a>

## Next pages — readiness_check / 110330101223 / 9

- [service.containers.readiness_check.exec_health_check](resources--workload--reference--group-016.md#canonical-2121212020120232-2033301021030021-0103132112101003-3221002101233112-2232113331320200-1122220030311221-0213111103312110-0220201122030101)
- [service.containers.readiness_check.http_health_check](resources--workload--reference--group-016.md#canonical-2002231301213013-0300131302213232-3012321112213220-0202230303103030-3332222333020013-3300112303113211-3301002230313323-1131103111202112)
- [service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2121212020120232-2033301021030021-0103132112101003-3221002101233112-2232113331320200-1122220030311221-0213111103312110-0220201122030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100303202320331-3001011300112331-0300322330013020-0312021330012332-1220213320112133-0303100112130303-2112331113202111-0203023200130110"></a>

## service.containers.readiness_check.exec_health_check — exec_health_check / 033302230121 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- service.containers.readiness_check.exec_health_check

<a id="canonical-1321101121232121-1103312120310203-2212301210110112-3110030233322322-2100131230313120-1031033120321113-3313313301221101-0030001300222031"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322110213223313-0222302220320000-1021110302001112-2203230203011323-1003223320111013-0031102303030023-0310000123222022-0203331310032023"></a>

## Direct properties — exec_health_check / 033302230121 / 3

<a id="canonical-2132313010310330-2020120310301200-2122122301010011-1132302331311232-0221300212212020-2103103111301012-0100301210101320-2320022131001203"></a>

<a id="canonical-3133101002233012-3320210301221320-0131000310113021-2031112311220031-0231133321312212-3012303112021110-1212311232231330-2121012210300233"></a>

## command property — exec_health_check / 033302230121 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1321212031020132-1311030123123232-1223231201133113-0230111203301100-3121123031023312-3302323333201131-3100213322032003-0301231300131001"></a>

## Next pages — exec_health_check / 033302230121 / 5

- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2002231301213013-0300131302213232-3012321112213220-0202230303103030-3332222333020013-3300112303113211-3301002230313323-1131103111202112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130013203111131-3123221202102213-1303022123111213-3222002323332031-2013213110121112-3322232010300312-3123231023230123-1212203332321111"></a>

## service.containers.readiness_check.http_health_check — http_health_check / 112230030111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- service.containers.readiness_check.http_health_check

<a id="canonical-3231000012130123-0033320101300313-2222103310303120-1012300031333110-3300130320223132-3233122232101233-2313013330033321-1000321011302311"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233000202101332-3131220332012222-2231003101202133-3012333113311131-2312212310213030-0121002003331232-1330312321203102-3102332110032103"></a>

## Direct properties — http_health_check / 112230030111 / 3

<a id="canonical-3122021310313112-2213132302200133-1211213220313100-0211213300203300-2303311121330122-3110222222202203-2303303121223320-3002021003031011"></a>

<a id="canonical-1010020203222222-2100220311321132-0203133202321200-1300003301030323-3103313323132220-1002201023222023-1033220301313002-2120033312002223"></a>

## headers property — http_health_check / 112230030111 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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

<a id="canonical-0102121230213020-2313111210301122-1133000231111000-2232221310310110-3200023120110122-0100300120021113-0112332330011210-1010320132020023"></a>

## host_header property — http_health_check / 112230030111 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

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

<a id="canonical-1203302111322132-2101233211020332-1022321032112002-2010223301333132-2000220131001330-0311330333220301-3021212230000220-2010122122111010"></a>

## path property — http_health_check / 112230030111 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0100301100001102-0321200213201203-0033123002230211-2112213010330202-0003321212203112-2133023303101111-3032110102101210-2103132200330222"></a>

## Next pages — http_health_check / 112230030111 / 7

- [service.containers.readiness_check.http_health_check.port](resources--workload--reference--group-016.md#canonical-0210321022333123-2321122000201010-2122130303010002-1212010310030122-3332222320112012-0100000032032322-1231111010210313-2311132203031320)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0210321022333123-2321122000201010-2122130303010002-1212010310030122-3332222320112012-0100000032032322-1231111010210313-2311132203031320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213233321001132-2313020130102300-0031022103300322-0200202310020221-2303212103130021-3211023310130320-2200202100100102-2102030310212333"></a>

## service.containers.readiness_check.http_health_check.port — port / 333322322102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- [service.containers.readiness_check.http_health_check](resources--workload--reference--group-016.md#canonical-2002231301213013-0300131302213232-3012321112213220-0202230303103030-3332222333020013-3300112303113211-3301002230313323-1131103111202112)
- service.containers.readiness_check.http_health_check.port

<a id="canonical-3023031203133323-1323220032103222-3300213102323020-1321210233103302-1311333201312310-2233111022222322-3010320110120203-2303023023201322"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121213132310031-0211110012010132-0332033020233121-0331131300200112-3201201033101002-3002131121033103-1022210000012130-2011212130031001"></a>

## Direct properties — port / 333322322102 / 3

<a id="canonical-3311113310000130-2013202023122010-2330101102302233-3002031320022233-3303322033020330-1212303212111111-3310211333131223-0331231030213011"></a>

<a id="canonical-1312000030320221-3212200122111011-3122032332231112-1203120301310321-0330020103113001-2231211310331021-0113311000030202-1203222032103320"></a>

## name property — port / 333322322102 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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

<a id="canonical-1103321302120320-3300300031202331-0113113310102323-2213300212123100-0230000032210210-0333232232220203-0132111033122030-1010130330210100"></a>

## num property — port / 333322322102 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0201201121103222-2021211012320220-0000201223322100-2331010000231333-2220233230032030-0120220113022022-1023001203023112-1122130221130322"></a>

## Next pages — port / 333322322102 / 6

- [service.containers.readiness_check.http_health_check](resources--workload--reference--group-016.md#canonical-2002231301213013-0300131302213232-3012321112213220-0202230303103030-3332222333020013-3300112303113211-3301002230313323-1131103111202112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230023133232220-0003303120110303-0303132310020203-1323112333302022-3201001310131202-0221220221011102-2323221200000011-0311100132322312"></a>

## service.containers.readiness_check.tcp_health_check — tcp_health_check / 103021002321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
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

<a id="canonical-3122321233103122-1211101003110320-1000210023033103-3203203000011323-0312103301300230-3022003322032203-0300001311332302-2323133021200020"></a>

## Direct properties — tcp_health_check / 103021002321 / 3

- [port](resources--workload--reference--group-016.md#canonical-2212311102201300-1321133023202333-1202331302220201-3300113313012031-3212022302113322-0323201201231011-1331022002230220-1131231303332213): complete subsection reference.

<a id="canonical-2310020000003233-2322220122102101-2033203102330221-1320100300332312-0311302230022012-0123300102111213-0201311030302212-0021010103311213"></a>

## Next pages — tcp_health_check / 103021002321 / 4

- [service.containers.readiness_check.tcp_health_check.port](resources--workload--reference--group-016.md#canonical-2212311102201300-1321133023202333-1202331302220201-3300113313012031-3212022302113322-0323201201231011-1331022002230220-1131231303332213)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2212311102201300-1321133023202333-1202331302220201-3300113313012031-3212022302113322-0323201201231011-1331022002230220-1131231303332213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333233210320003-2301100100233003-0221210111032100-0203232132233022-1022232103131331-2002010312203223-0232113212100322-3011020222103203"></a>

## service.containers.readiness_check.tcp_health_check.port — port / 320130230331 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122)
- [service.containers.readiness_check](resources--workload--reference--group-016.md#canonical-2030201031231200-3332330302130030-0232113021223033-1003123031020103-3030133132122013-0031120200130120-1030220221132302-2102321302231233)
- [service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323)
- service.containers.readiness_check.tcp_health_check.port

<a id="canonical-0123122132011210-0313023320103313-2112302113011303-3230023212120223-2023003312200112-1101133023312213-3332323032100232-2232311020001320"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210231310032112-0011330232233220-3000200013232133-2231301232000110-3002311002211233-2303230011131221-1231113330021222-1311131333032231"></a>

## Direct properties — port / 320130230331 / 3

<a id="canonical-1213333210320330-0121211221123110-0301312311101012-1322320232322330-2012100232233312-3013231123302002-0233002001311221-3123130223203131"></a>

<a id="canonical-1033332013131101-3123003221230123-3310110110222201-2222130223322313-0331212222231312-1020003123222232-1130332111200222-0033311113331230"></a>

## name property — port / 320130230331 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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

<a id="canonical-2321330010132310-1032333120213232-3001011022110210-2123021013120031-0231323001333111-0220222102233021-1010320211213233-1103000030122312"></a>

## num property — port / 320130230331 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1032100203303130-2302100100030022-2021330120103101-1233323022022333-1313002033021120-2010220012213011-0201023212311223-2211301123030301"></a>

## Next pages — port / 320130230331 / 6

- [service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032010020312200-2121131321210330-2323110003203001-2023303033223232-0120033020132022-2331023002032132-1001211311110212-1301312100113233"></a>

## service.deploy_options — deploy_options / 102123100031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.deploy_options

<a id="canonical-2003022010122311-3321021022020131-2103000013201023-1222302021013313-3001303103030311-1332132113321232-2300111213130322-3321232233101331"></a>

Type: `"object"`. single nested block, Optional.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_res",
    "default_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_re_sites",
    "deploy_re_virtual_sites")}
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
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

Terraform syntax:

```terraform
deploy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3001321102203103-3112312210003003-3332002220001013-1132312330333310-1330220111220313-3103103302331210-1131331033100103-3313132002032222"></a>

## Direct properties — deploy_options / 102123100031 / 3

- [all_res](resources--workload--reference--group-016.md#canonical-3320122331212021-2103131122203112-1021030123303133-1113130130300120-1332001011321111-1021203031103220-1133320202311232-2230100023300313): complete subsection reference.

- [default_virtual_sites](resources--workload--reference--group-016.md#canonical-0123021220301230-0122112311122001-3213103300231202-3110232110313113-2030320231100102-1111033100101031-3212301103121201-1001132210103332): complete subsection reference.

- [deploy_ce_sites](resources--workload--reference--group-016.md#canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213): complete subsection reference.

- [deploy_re_sites](resources--workload--reference--group-016.md#canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202): complete subsection reference.

<a id="canonical-1222133311232022-2322133202212113-1303103303221231-1111233132320321-0033302322110213-1102111200212132-1000312032031232-1312230113201111"></a>

## Next pages — deploy_options / 102123100031 / 4

- [service.deploy_options.all_res](resources--workload--reference--group-016.md#canonical-3320122331212021-2103131122203112-1021030123303133-1113130130300120-1332001011321111-1021203031103220-1133320202311232-2230100023300313)
- [service.deploy_options.default_virtual_sites](resources--workload--reference--group-016.md#canonical-0123021220301230-0122112311122001-3213103300231202-3110232110313113-2030320231100102-1111033100101031-3212301103121201-1001132210103332)
- [service.deploy_options.deploy_ce_sites](resources--workload--reference--group-016.md#canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323)
- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213)
- [service.deploy_options.deploy_re_sites](resources--workload--reference--group-016.md#canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030)
- [service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3320122331212021-2103131122203112-1021030123303133-1113130130300120-1332001011321111-1021203031103220-1133320202311232-2230100023300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030213120100102-1123310202322100-2110120001031200-2021301121203021-0230011001230220-0132113322000221-3122102120310031-2332103212122232"></a>

## service.deploy_options.all_res — all_res / 030223301023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.all_res

<a id="canonical-1321313220330303-2030200020321220-2201323120232121-3001200001302102-2212210330130213-3210010332323111-0030302222202312-2323123020220300"></a>

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
all_res = {}
```

<a id="canonical-0201223322100230-1332203101211230-0030230100232120-3010032231112022-1302301020322003-0133302333223213-0311322321333030-3033033011332031"></a>

## Direct properties — all_res / 030223301023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130232322100323-1222022100033121-1110322101221331-0313210322020111-3123111133000011-1310020223133120-3220101111300021-2100221200212001"></a>

## Next pages — all_res / 030223301023 / 4

- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0123021220301230-0122112311122001-3213103300231202-3110232110313113-2030320231100102-1111033100101031-3212301103121201-1001132210103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022212103021000-1320011112032321-1033233220002011-1200303032003102-0201311333203133-2333311301130023-3220032332201221-0113210333121200"></a>

## service.deploy_options.default_virtual_sites — default_virtual_sites / 212301003201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.default_virtual_sites

<a id="canonical-1321212010121323-3300132300321312-1223311311002222-1012312001300000-0032330033120033-0003032322003211-3301133200012103-3003112022230201"></a>

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
default_virtual_sites = {}
```

<a id="canonical-2320202131003300-0231103312202031-1320321322313130-1013313231321112-2000313222200222-3320201011312323-3002310130313231-0133300223200013"></a>

## Direct properties — default_virtual_sites / 212301003201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012102302212201-1330231221033233-1200330313313031-3131321301110300-1200323233211122-2002233221211133-2003131133020113-2310001312200203"></a>

## Next pages — default_virtual_sites / 212301003201 / 4

- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100132213233323-0130321121321113-1131332022302000-3020132130222003-0322301130021303-3331121320012332-0102010200320123-3332032132121011"></a>

## service.deploy_options.deploy_ce_sites — deploy_ce_sites / 100220313201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_ce_sites

<a id="canonical-2121110312021123-3303032202132020-0101132103323032-2213131023102202-2332222102323102-3002131213322233-0223101120331133-2311311023101112"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121130032201222-1223012221212321-1022021012301211-2033323032321200-1222023020010201-0011133103232331-1103220023001320-3332221221030311"></a>

## Direct properties — deploy_ce_sites / 100220313201 / 3

- [site](resources--workload--reference--group-016.md#canonical-3101000002123000-1233232311011210-3101322210100023-3232310031032321-1123000020111333-3201031203033230-1312221132032301-1322320221213110): complete subsection reference.

<a id="canonical-0110313100103201-3030311312333230-1033300333301010-0130101013122000-0002133213001130-2201010221112023-2301131311113020-0211310222221131"></a>

## Next pages — deploy_ce_sites / 100220313201 / 4

- [service.deploy_options.deploy_ce_sites.site](resources--workload--reference--group-016.md#canonical-3101000002123000-1233232311011210-3101322210100023-3232310031032321-1123000020111333-3201031203033230-1312221132032301-1322320221213110)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3101000002123000-1233232311011210-3101322210100023-3232310031032321-1123000020111333-3201031203033230-1312221132032301-1322320221213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313311032300130-3332302230110300-3130310230120021-1333212133010221-1330132312300320-0121201301222223-2200210101230000-1101120200020110"></a>

## service.deploy_options.deploy_ce_sites.site — site / 202222121211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_ce_sites](resources--workload--reference--group-016.md#canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323)
- service.deploy_options.deploy_ce_sites.site

<a id="canonical-3022112202330230-2323132320333020-3310232310031222-1033022200310011-3011313033220232-2001311210030323-2320203202110223-0313323113111323"></a>

Type: `"object"`. list nested block, Optional.

Which customer sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3201033302301110-0211110310031222-3312111013111122-0021311033211220-1133211132323000-1131131310313200-1022203223121003-2232010113321123"></a>

## Direct properties — site / 202222121211 / 3

<a id="canonical-1300310032022010-2301202212203020-2001231010011123-1130311313111123-3023122100133113-0202033232311011-2320322130222230-3121111022220122"></a>

<a id="canonical-2013223021003322-1011002220210133-2210000102212112-2111223132311331-3122303122032213-0311033233222101-1222223103301312-3213211102223012"></a>

## name property — site / 202222121211 / 4

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

<a id="canonical-3003221210321002-3132210000331333-0301120321222221-0012322302221122-3101323001000121-3012003303221313-2100130021032213-2302130011320031"></a>

<a id="canonical-0020011012301313-1002112011213310-3111101210312011-3023201321303102-2231213102220300-0001033323031310-0213330013121020-0301011303033302"></a>

## namespace property — site / 202222121211 / 5

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

<a id="canonical-3011221310022020-3313233302330111-0320032032330032-2201133333331013-1300310231131202-0321230313021303-3022313222103130-2223112210221323"></a>

<a id="canonical-1212202310010033-3302223131133220-2133113312303231-0213213102232132-1222211013113010-3300203021121201-1220100212110202-1330103203200132"></a>

## tenant property — site / 202222121211 / 6

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

<a id="canonical-3022303133310013-1122313013211122-1313111233313030-2312112011120221-0312333103130312-3122033211101131-1032332013203033-2320123232111131"></a>

## Next pages — site / 202222121211 / 7

- [service.deploy_options.deploy_ce_sites](resources--workload--reference--group-016.md#canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130032131113101-0332330101230130-0003133223002021-2123221330201321-1031221232113002-3213302221211332-2320132231232200-1222011012313220"></a>

## service.deploy_options.deploy_ce_virtual_sites — deploy_ce_virtual_sites / 023102220013 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-3220020110313002-1101331223200023-0333220300121021-0302330110232100-2122233122230312-1312201231103021-2201020012022111-1030100333101110"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Customer virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
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
deploy_ce_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200232010222223-2310031221120131-2130302021200200-2033103220013132-3200310120322203-0311011002232130-3013301232322222-0311011302130231"></a>

## Direct properties — deploy_ce_virtual_sites / 023102220013 / 3

- [virtual_site](resources--workload--reference--group-016.md#canonical-1201201122302303-1001333103123313-0130101233033100-3200113022111130-0213333133332302-3100202312211131-2120213320312021-0332232100003110): complete subsection reference.

<a id="canonical-3011121310221222-0320202100233102-2310002322112020-2203312202133322-3010001221112321-3231020331220201-3023302303123122-2321000032121131"></a>

## Next pages — deploy_ce_virtual_sites / 023102220013 / 4

- [service.deploy_options.deploy_ce_virtual_sites.virtual_site](resources--workload--reference--group-016.md#canonical-1201201122302303-1001333103123313-0130101233033100-3200113022111130-0213333133332302-3100202312211131-2120213320312021-0332232100003110)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1201201122302303-1001333103123313-0130101233033100-3200113022111130-0213333133332302-3100202312211131-2120213320312021-0332232100003110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203130212000130-2120203132100331-1210200200332321-0303032021232212-1012203121332001-1211203213330303-3030310000002322-3103312221022202"></a>

## service.deploy_options.deploy_ce_virtual_sites.virtual_site — virtual_site / 201220333303 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213)
- service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-2132123023233221-2023233213131021-0201022213110220-0132110011200013-1223023220112011-2013133033221121-1002231221113201-2110232001223123"></a>

Type: `"object"`. list nested block, Optional.

Which customer virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1322212312022122-2222313011202200-1313231303332320-1232213200211213-1330332320322120-1312102100102231-1222022221302130-3103300202113203"></a>

## Direct properties — virtual_site / 201220333303 / 3

<a id="canonical-2223103130011133-1210222212312232-0321320102102323-1003332100131223-1122230210300013-2301212121100112-3103111302102121-0132122121002222"></a>

<a id="canonical-2221333022101020-2122330101123230-2031301123003311-2301211100323222-0122233312023202-3231201310112203-2133103330201223-1233033022011122"></a>

## name property — virtual_site / 201220333303 / 4

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

<a id="canonical-0000201200101033-3311031231220200-0120211203211321-3300131021312102-0201312300201321-2303333033013303-1333203021032001-0321001210021002"></a>

<a id="canonical-0120232212102221-1000020030021103-0300023111132223-3213210122333011-1032022031302321-1022233213101302-2101202113323031-1001032102103030"></a>

## namespace property — virtual_site / 201220333303 / 5

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

<a id="canonical-2232202030332313-1022232010113303-3301030202122221-0211320103323103-2101201130030030-3231331220013202-2323221011320303-2101212100011212"></a>

<a id="canonical-0220121212002001-0330002221013031-2233032232223222-1031132032100021-0103132020102323-1330212333311323-2300130202333313-2123332313333012"></a>

## tenant property — virtual_site / 201220333303 / 6

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

<a id="canonical-3012211233101322-1022002101120122-0222002233232321-3023111010202212-1032203221003012-2113320010123032-0003310122100210-0323210010323202"></a>

## Next pages — virtual_site / 201220333303 / 7

- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-016.md#canonical-1022033030230100-2023233310231233-0101322231213121-1200031301010320-1010203312213312-0032021033302132-2120223320233022-1301331000021213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003333121220210-2312100313323110-1310132321033301-1320220000133032-3322122211300210-0130002130121003-2222100002131201-3201331010100313"></a>

## service.deploy_options.deploy_re_sites — deploy_re_sites / 233321102303 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_re_sites

<a id="canonical-3013322233213022-2310103223230200-1303031300003322-3323012231322013-0211003110301321-2220020221322210-1330311000100132-3332112313001333"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011221020112233-0203213233011001-3110023203310203-2100303222110133-3213221220233322-2212301112001300-2230011311011010-1011113301130113"></a>

## Direct properties — deploy_re_sites / 233321102303 / 3

- [site](resources--workload--reference--group-016.md#canonical-2033131030312023-2200112222221310-2222232223012011-0200121201203231-3033103333011032-0331022200133101-2230032301132121-2012311230312312): complete subsection reference.

<a id="canonical-0333131020113113-1313320331103011-1221011110212311-3320201112120113-0220103030321130-1001032232302130-0011033301103213-3322322211010120"></a>

## Next pages — deploy_re_sites / 233321102303 / 4

- [service.deploy_options.deploy_re_sites.site](resources--workload--reference--group-016.md#canonical-2033131030312023-2200112222221310-2222232223012011-0200121201203231-3033103333011032-0331022200133101-2230032301132121-2012311230312312)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2033131030312023-2200112222221310-2222232223012011-0200121201203231-3033103333011032-0331022200133101-2230032301132121-2012311230312312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122230130031210-0202201133123223-0112102301332032-3300003022200103-3211000312112231-2122212032302203-0332103103000313-1122012131210103"></a>

## service.deploy_options.deploy_re_sites.site — site / 102012313011 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_re_sites](resources--workload--reference--group-016.md#canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030)
- service.deploy_options.deploy_re_sites.site

<a id="canonical-3321022030130120-1002110101331302-0100203320122111-3111000102200000-3322221112301121-2113110111332210-2101301020111222-2030221002021221"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1330312133301132-1131100121103122-3032332110333110-2230123212330231-1332233112310032-2132030220302303-3230313100302013-1000231023222101"></a>

## Direct properties — site / 102012313011 / 3

<a id="canonical-0113311010313211-1013213202321231-2013023011230210-3232202122330300-1030302312320332-1021322213112303-3222110132133210-2111011033323121"></a>

<a id="canonical-2112031220312221-3300233012112203-3000210103101200-1133110032230102-0113102111011112-3131130033032230-0100121010321000-3202001300110302"></a>

## name property — site / 102012313011 / 4

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

<a id="canonical-0203321021210101-3203000310331111-2011311122130021-1130111322332111-3331103212230111-3212320213200222-2013203010231030-1303223120313310"></a>

<a id="canonical-2223303112202323-3311003320030311-0003031002113010-1122202000223112-1033002001332011-1222333331231233-2312011120100132-0122130112313022"></a>

## namespace property — site / 102012313011 / 5

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

<a id="canonical-1123200303200102-0203230322213110-1021100202020233-1001100222232112-1210213313022320-0033320323321222-1021320322232211-2223322303212031"></a>

<a id="canonical-2211210103312110-0311111030322003-0213011222232321-0330223310312032-3101221111203013-0333203233231021-3300211333113103-0203022221302101"></a>

## tenant property — site / 102012313011 / 6

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

<a id="canonical-0211320010320223-3113302301101231-3112222320011332-1010322023121322-1221103213203202-0021232220210130-3020102112101332-3301222130212321"></a>

## Next pages — site / 102012313011 / 7

- [service.deploy_options.deploy_re_sites](resources--workload--reference--group-016.md#canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211212120232321-2012132112021001-3120233010032313-1301102120113013-1233002022111121-2000232133233113-0312323013003121-3211330110332323"></a>

## service.deploy_options.deploy_re_virtual_sites — deploy_re_virtual_sites / 110033312112 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- service.deploy_options.deploy_re_virtual_sites

<a id="canonical-0222222021111203-1231331110311132-3303120220302320-0011221210101320-0230110102300222-1122020302302120-2123131030030213-3320230000331223"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
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
deploy_re_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220132203323230-2000321303030310-3131221331103113-3313133132232010-1001211221002021-3322233121302003-2311002323321333-1332111232000300"></a>

## Direct properties — deploy_re_virtual_sites / 110033312112 / 3

- [virtual_site](resources--workload--reference--group-016.md#canonical-2322220331033200-1103303020112230-3013011001212213-3010301232101123-2013313020321222-3011032223303110-2121123030131323-0231033133133200): complete subsection reference.

<a id="canonical-0222233310230330-0222121030123100-2132013122033132-0311322211301133-1231012130232331-2103131133223103-0330033301033213-1001300102010311"></a>

## Next pages — deploy_re_virtual_sites / 110033312112 / 4

- [service.deploy_options.deploy_re_virtual_sites.virtual_site](resources--workload--reference--group-016.md#canonical-2322220331033200-1103303020112230-3013011001212213-3010301232101123-2013313020321222-3011032223303110-2121123030131323-0231033133133200)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2322220331033200-1103303020112230-3013011001212213-3010301232101123-2013313020321222-3011032223303110-2121123030131323-0231033133133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203230220022201-3320113020010010-2331133010031312-3111323132000023-0031333101133213-0032233221032313-0032313130030310-1123332012231230"></a>

## service.deploy_options.deploy_re_virtual_sites.virtual_site — virtual_site / 303313223201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003)
- [service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202)
- service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-3122121120213322-1321020313301213-2303020310132333-0102230231021201-3313011220013121-1200331201000310-3033131201030031-2313222022013023"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3331000110330320-2212123201133303-3012303333331312-0201132000332012-2230233121103232-2120223023112020-0313303120223310-3321122301222023"></a>

## Direct properties — virtual_site / 303313223201 / 3

<a id="canonical-3003323132012212-0132303323030003-0203200231311011-2001333132303030-3012310123013020-2022103010200311-0301010021211131-2122212022113323"></a>

<a id="canonical-2323132011131313-2201013313320302-0023220030312301-3232133112233123-3211201313002202-1211232301323122-3323031310323020-3211002201212313"></a>

## name property — virtual_site / 303313223201 / 4

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

<a id="canonical-2311203021302013-3203110322311321-1202022013303300-2311232121301331-1232313103303311-0202321330311113-1113230332200201-2123022230132112"></a>

<a id="canonical-1020230213021203-1103310111211010-2331021100011102-0011310213321313-2022110001023022-2310123000232232-2200010310101220-1200311213303033"></a>

## namespace property — virtual_site / 303313223201 / 5

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

<a id="canonical-1222330022103120-2202313202211022-3203020023021123-3210330322210233-3121311001230131-3330021303023223-2212310300202110-2013131221031330"></a>

<a id="canonical-0123112301312100-0001002111310200-3033012012000221-2033320220310131-2033132221033003-1130031332012000-3203021300212313-0213222112023310"></a>

## tenant property — virtual_site / 303313223201 / 6

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

<a id="canonical-3333233231203211-2313220313012120-2310113301031201-3212313111312210-0111013001301130-0231313200111320-2203110233121121-2323231022121111"></a>

## Next pages — virtual_site / 303313223201 / 7

- [service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-016.md#canonical-2022112111312113-1223330120022302-3112030012223101-0031101323310000-0210111313202312-1031011120212112-1232312211233332-2011011322131202)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2112203323032023-1022033211220013-2331111010311212-0200221020302323-2123121101110311-2211200012322311-2122132200232301-3022222221010211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231200210330000-2132031003102013-1012201323222301-0220022323021201-0133300213133103-3323022003013310-3030023012330010-2131120132031323"></a>

## service.scale_to_zero — scale_to_zero / 301013133030 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.scale_to_zero

<a id="canonical-3012300321212121-1132001221112031-2333022003222023-3111011130310012-3031202110320300-3011220322003230-3302003002322201-0011302201032002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for scale to zero.

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
scale_to_zero = {}
```

<a id="canonical-3110333211013302-1222211133011322-2131322120230210-1100203233323110-0132010113010021-3012221020030313-0320130203103012-3310230211100133"></a>

## Direct properties — scale_to_zero / 301013133030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010013131110101-2333000233330002-2132223002201100-0012222223010131-2020333211310112-1231330300302322-3013023013201322-3221302333122200"></a>

## Next pages — scale_to_zero / 301013133030 / 4

- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020220300310201-3011211220000330-0310111033220300-3311111121313320-0122201122330103-3210212320301203-1111030223133023-3223210232120333"></a>

## service.volumes — volumes / 312230021100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.volumes

<a id="canonical-0323112231212112-2031033023310330-0111121103130221-0300012223232011-0303013233030011-1113121231003200-3012212103230331-1211323311213333"></a>

Type: `"object"`. list nested block, Optional.

Volumes. Volumes for the service.

Upstream description:

Volumes for the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path"),
  validators.ConflictingListObjectAttributes("empty_dir",
    "persistent_volume"),
  validators.ConflictingListObjectAttributes("host_path",
    "persistent_volume")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
volumes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033310011033013-2301201203122103-0203310313313203-2331002122122031-3123201010031322-1122120031220211-0110130021101210-1210001310213012"></a>

## Direct properties — volumes / 312230021100 / 3

- [empty_dir](resources--workload--reference--group-016.md#canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000): complete subsection reference.

- [host_path](resources--workload--reference--group-016.md#canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223): complete subsection reference.

<a id="canonical-0210321122110332-3332130230330223-0310321033222331-1120102210232313-1113320201201213-1132320030221313-0033130303003012-0301121201203300"></a>

<a id="canonical-3003030211001121-3102003103130323-3021230131323123-2111220102230333-3111301120110031-3330122310021202-1011200221131122-3100103300200123"></a>

## name property — volumes / 312230021100 / 4

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](resources--workload--reference--group-016.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331): complete subsection reference.

<a id="canonical-2312021302131231-0022033110213300-0121021020100221-0223323031333300-0013333220320012-1020331233021333-3202012312230200-2123221303200201"></a>

## Next pages — volumes / 312230021100 / 5

- [service.volumes.empty_dir](resources--workload--reference--group-016.md#canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000)
- [service.volumes.host_path](resources--workload--reference--group-016.md#canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001201103000103-0203100000113101-0212022211332012-2101001231333203-3011232200133102-0200323032331200-3301303002100313-0220200113001312"></a>

## service.volumes.empty_dir — empty_dir / 120012213222 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- service.volumes.empty_dir

<a id="canonical-0013300133100100-1113110123121321-3111203122110000-3001230320302321-1033333323033011-2102333222220223-0221203100323332-1123020123001331"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("size_limit")}
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
empty_dir {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203312030221323-3031332010313120-0133031110002023-3122212222300333-2031013112311132-0313100200130102-0322200332301220-2203330003331100"></a>

## Direct properties — empty_dir / 120012213222 / 3

- [mount](resources--workload--reference--group-016.md#canonical-3202131021131212-0101213023301012-0231220211323020-2331102001300001-1020233213033112-1212100032103200-1202110312132003-2211301123313232): complete subsection reference.

<a id="canonical-0132332000113213-1120211203132332-3200303132331322-0001321132202203-2231132001003331-1312030201123212-2213103232013103-1231211031322233"></a>

<a id="canonical-1333212003100223-2111030131213330-1331113101003231-3031331102321331-1320001322210111-2330200222120031-3220221001333123-1230122322321310"></a>

## size_limit property — empty_dir / 120012213222 / 4

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1131230311120132-1202011333311202-0220021213011120-2221133233120023-0110102011321003-1120330200303312-2301102120113022-0212013313111020"></a>

## Next pages — empty_dir / 120012213222 / 5

- [service.volumes.empty_dir.mount](resources--workload--reference--group-016.md#canonical-3202131021131212-0101213023301012-0231220211323020-2331102001300001-1020233213033112-1212100032103200-1202110312132003-2211301123313232)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3202131021131212-0101213023301012-0231220211323020-2331102001300001-1020233213033112-1212100032103200-1202110312132003-2211301123313232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112032131032123-1221321212011002-3131030320321002-0032330200210311-1333132023122320-2201322233102220-1033002202303300-1012130013222231"></a>

## service.volumes.empty_dir.mount — mount / 230330103233 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.empty_dir](resources--workload--reference--group-016.md#canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000)
- service.volumes.empty_dir.mount

<a id="canonical-3200110313320120-1301012232022302-3002200123222132-1113212133033223-2222012030122222-0331101330013120-3213301223312132-3123100222010012"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131321003220330-3002022202211121-2131221323301023-2112120013233102-3303302123111122-3232311011232010-3333022311101033-3300132030232021"></a>

## Direct properties — mount / 230330103233 / 3

<a id="canonical-0011311212322030-3000213100333003-0303013013110321-1020212111032132-0300121132203120-2022133203301113-1303033232220120-3332212301120201"></a>

<a id="canonical-2320133203113010-2221201103323120-3022232321232232-3121211222013030-3332211022302302-1321002033030303-1013310033030020-0022012030122201"></a>

## mode property — mount / 230330103233 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2112223001233030-1032021033333131-0122001321010212-0230122121002230-3102031231203323-2310221123231022-0333031002333201-2120311101020022"></a>

<a id="canonical-1221011012202301-2133333301003030-0323123032003313-3232112312112332-1112011103202001-2021123013102031-2310013223332220-0221331000313010"></a>

## mount_path property — mount / 230330103233 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-1223032003000302-0200111301032101-3203103211120130-2211301313111220-3233012131033321-2013122331220133-0110102211321033-1213130323033000"></a>

<a id="canonical-0223111100101113-3310233122233101-3133103311213012-2110220311221130-1233232203231201-1221023320010201-0322313133133101-3102332032130010"></a>

## sub_path property — mount / 230330103233 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2331200032223101-1101310221030230-3233310102022310-1320103020112330-2332210130311030-1110001302203213-3312000230300003-3301223302001033"></a>

## Next pages — mount / 230330103233 / 7

- [service.volumes.empty_dir](resources--workload--reference--group-016.md#canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331132311012233-2010132130010233-2032313023322010-3200220022111132-3221320013013301-0022013020301113-2202212120232113-0212313013330112"></a>

## service.volumes.host_path — host_path / 201121200010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- service.volumes.host_path

<a id="canonical-0102110212031231-2121203200102230-0133203021023222-1310331100120120-1121032013330202-2123121310133130-0311210032303203-0023311221310001"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a host mapped path into the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
host_path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200122132122033-1132333200123033-3303000203221221-1233222121113120-0022032323032110-2213100101021021-3301211212031131-3203001020021331"></a>

## Direct properties — host_path / 201121200010 / 3

- [mount](resources--workload--reference--group-016.md#canonical-1303131310232102-2011333101230102-1121333013101031-3102011101122303-0320111112321001-1102313022211132-2210303010130031-0203210032311333): complete subsection reference.

<a id="canonical-1210301313203011-3113123120033122-2223131110021130-3031310011100010-3131233113333112-2223210232113201-2302323113122312-1223321030231303"></a>

<a id="canonical-3201111011203232-2102221223112213-1321012120300021-1301133332320222-0130213100322012-0002333123223013-2020230223131123-3031102130212213"></a>

## path property — host_path / 201121200010 / 4

Type: `"string"`. Optional.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

<a id="canonical-2331003200321333-3032221333120330-1332103302333231-1202232312033210-2032101033032100-0032232321010202-1313130100220222-1211100220121332"></a>

## Next pages — host_path / 201121200010 / 5

- [service.volumes.host_path.mount](resources--workload--reference--group-016.md#canonical-1303131310232102-2011333101230102-1121333013101031-3102011101122303-0320111112321001-1102313022211132-2210303010130031-0203210032311333)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1303131310232102-2011333101230102-1121333013101031-3102011101122303-0320111112321001-1102313022211132-2210303010130031-0203210032311333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303013000202133-3013313222002303-2303103323113012-1201130223130100-3303201200111020-3320300033210222-1212020231002002-3022132221012010"></a>

## service.volumes.host_path.mount — mount / 312002200322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.host_path](resources--workload--reference--group-016.md#canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223)
- service.volumes.host_path.mount

<a id="canonical-0321101100023131-2010212020031132-3110311213201330-2010122221122013-1020333102322302-1102230121332220-2301312003323333-2013131022001013"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333330013130102-3013022102031020-1301330033221232-3033021313201020-2220301320001312-0321102100131011-0003202132203130-3333203310011223"></a>

## Direct properties — mount / 312002200322 / 3

<a id="canonical-1110132323310221-2311103023310013-0120310210033123-1132303132123210-0111323323020233-2222122031330032-2101113302001303-2030123012011332"></a>

<a id="canonical-3213111221100131-0103130322120131-3123010020303123-1103000220131101-3300202013230023-0010212003120211-1211031233033323-1133133022302221"></a>

## mode property — mount / 312002200322 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0003210232113002-1020033232313322-0323303012133323-3100003101000032-1232211101323133-3300022330232013-1223122311330202-0233310110031311"></a>

<a id="canonical-3210303032230210-3121031023233322-2013020100023330-3320103331202213-3102101203012110-1033102131312023-0032220210132311-3311133112012132"></a>

## mount_path property — mount / 312002200322 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-3221101333020313-2033123201020102-0011022230030232-3120301332300202-3131220322003001-0022032201321010-3201230030132323-0301120331002100"></a>

<a id="canonical-3023001011111012-1032322200130331-1320203122133102-1002122312322312-1101103001012212-2212022121210333-0002311010100200-1222100202120312"></a>

## sub_path property — mount / 312002200322 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0121331320010223-0122010330320121-1030111130233330-0330033233033233-0323013220000023-3220332320222120-2221002310112222-2133033011213020"></a>

## Next pages — mount / 312002200322 / 7

- [service.volumes.host_path](resources--workload--reference--group-016.md#canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330012312210333-3203020323322131-2033112023212311-1101300233311032-2231030121210233-0313120202210013-3330123202223101-1012033121122103"></a>

## service.volumes.persistent_volume — persistent_volume / 033001321032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- service.volumes.persistent_volume

<a id="canonical-0331030333110000-3311030122030203-3313001232120320-1301111312302322-2101030020330020-0103202221303122-0013110333330223-1110000022313101"></a>

Type: `"object"`. single nested block, Optional.

Volume containing the Persistent Storage for the workload.

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
persistent_volume {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123310202311223-2301222033011033-0020330130301012-0103321231010011-1113223312100231-1212010113202222-3101103021110233-3010300301133131"></a>

## Direct properties — persistent_volume / 033001321032 / 3

- [mount](resources--workload--reference--group-016.md#canonical-0230211303132020-2130130313320130-3020231002113003-1021322300103102-0220013320231100-0021103211212131-2122030310213311-3112100123213012): complete subsection reference.

- [storage](resources--workload--reference--group-016.md#canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312): complete subsection reference.

<a id="canonical-1101200311112101-1310013122031031-1321230032021102-0102032213320221-2013112020113012-2002313003101332-1000001311311320-1113332232230131"></a>

## Next pages — persistent_volume / 033001321032 / 4

- [service.volumes.persistent_volume.mount](resources--workload--reference--group-016.md#canonical-0230211303132020-2130130313320130-3020231002113003-1021322300103102-0220013320231100-0021103211212131-2122030310213311-3112100123213012)
- [service.volumes.persistent_volume.storage](resources--workload--reference--group-016.md#canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0230211303132020-2130130313320130-3020231002113003-1021322300103102-0220013320231100-0021103211212131-2122030310213311-3112100123213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021001311013321-0220223232231302-2131220233103231-1032210100033013-2013333302121302-2131101310311302-0200013030121211-1202302030332100"></a>

## service.volumes.persistent_volume.mount — mount / 121123303121 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- service.volumes.persistent_volume.mount

<a id="canonical-2201312230202120-2212303012330022-1312101133332121-3122030020131131-1202230120332130-1100331022032002-3120132023213113-1313210313131102"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323013230303312-2120021212131223-3210102113323220-2101220102323213-3213310102001321-3313123330030320-3020230303113121-1021013030232122"></a>

## Direct properties — mount / 121123303121 / 3

<a id="canonical-1203302302020000-3033220120113321-3133103221133201-3021122333031201-0231331333201031-1232322312311031-2032022003311011-3303211200333223"></a>

<a id="canonical-2220001320000210-1121321320022303-0301110333300232-0322222022123203-1003032330203313-1131221330100010-3313313022330203-3330220310230220"></a>

## mode property — mount / 121123303121 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2112303321133312-1310333020010002-0221202303110032-0323311303322323-1021122223100133-2322203221131220-0220112023131300-3221121232002112"></a>

<a id="canonical-0232121123212110-1321212113201033-2020010002032110-0030101321110112-1022211300002200-3020121122331302-2120223133003231-0210131313030013"></a>

## mount_path property — mount / 121123303121 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-0213021022230223-3203233101132120-0303222112130111-0133122020332301-0200103200200012-3210000122220203-2320323122302012-3013103010201220"></a>

<a id="canonical-0022123013131311-1012300221222013-3321332232111212-1221020023002013-3320113221231333-1201311021200223-2133020102102201-3221030230112201"></a>

## sub_path property — mount / 121123303121 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1013110222310113-3020123331320030-1020321013002131-0302100332233323-2230320000301300-3121310222312021-1200230030132200-0202001310322233"></a>

## Next pages — mount / 121123303121 / 7

- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220010331212122-1320100021313213-0201031031033021-2223331133101230-2030023101001002-0211011113233020-1210113121330101-1120332231010003"></a>

## service.volumes.persistent_volume.storage — storage / 200233100002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- service.volumes.persistent_volume.storage

<a id="canonical-1133211230011302-3231213011010102-1032132310232121-1213303303232213-1333211102102232-1002333020110033-3301332002031300-0233131221122002"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
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
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330002010122300-0030300022021223-0233000233131022-0230210231013100-2102221301002233-2111120321102001-2020220002321221-1013113313111332"></a>

## Direct properties — storage / 200233100002 / 3

<a id="canonical-0121011120331330-2111121210033102-3030110210222021-1203100330132200-0213220202303201-2002030323211311-2203211122213303-0201330330102102"></a>

<a id="canonical-2332010332200031-2002033020231300-2202022132221231-3200303032223211-3232021201111320-0220021220132022-3202132031332201-1210232102033210"></a>

## access_mode property — storage / 200233100002 / 4

Type: `"string"`. Optional.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2101130010203102-2302133001230232-2210321033131213-2300233012103302-2320013331130211-0032110001132103-1213033213223033-3303230333220201"></a>

<a id="canonical-3111302033230231-0131303010323200-1103020322201022-3121100320032130-0221013122321231-3311223233232233-0222033302033232-3112102200033001"></a>

## class_name property — storage / 200233100002 / 5

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](resources--workload--reference--group-016.md#canonical-3030021232102132-3032123220212030-1212020002332312-1203311122221322-2301311133233201-1131300212232312-1331331120023212-0103123233033101): complete subsection reference.

<a id="canonical-1000011321211123-1330211120011013-0111212123123010-0330232300322032-2303101101322300-3323300203010203-1202302321122313-3311211001020000"></a>

<a id="canonical-2203202013321123-0112230302313203-2031131111332003-0003223202200222-3233110121202322-1033101003302230-1131310312313120-3131120333013010"></a>

## storage_size property — storage / 200233100002 / 6

Type: `"number"`. Optional.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2102311010011302-3013320213322302-1222220112011122-3221232001020112-1112031211110200-0020022313201330-1302331322210120-3133220232303101"></a>

## Next pages — storage / 200233100002 / 7

- [service.volumes.persistent_volume.storage.default](resources--workload--reference--group-016.md#canonical-3030021232102132-3032123220212030-1212020002332312-1203311122221322-2301311133233201-1131300212232312-1331331120023212-0103123233033101)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3030021232102132-3032123220212030-1212020002332312-1203311122221322-2301311133233201-1131300212232312-1331331120023212-0103123233033101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122210322330110-1222312221330230-3101332003232333-3121200223033312-1311231002322113-3311220201032021-2033211333313330-1331010211230301"></a>

## service.volumes.persistent_volume.storage.default — default / 132002001201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003)
- [service.volumes.persistent_volume](resources--workload--reference--group-016.md#canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331)
- [service.volumes.persistent_volume.storage](resources--workload--reference--group-016.md#canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312)
- service.volumes.persistent_volume.storage.default

<a id="canonical-3232323111203121-0022330230132021-2122021001122201-1311212300230212-3203320002100233-2310222010011212-3122323001311033-2313002010023222"></a>

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
default = {}
```

<a id="canonical-3211320111311120-0100301233212112-3121301101021130-2301131221300302-2220333013000313-3222302320010230-0210000300301100-0023222201321222"></a>

## Direct properties — default / 132002001201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210302023202311-1312023010002012-1113023031111210-0231103113333232-2301301223000333-2313022031021323-2100313103012131-1331013213312121"></a>

## Next pages — default / 132002001201 / 4

- [service.volumes.persistent_volume.storage](resources--workload--reference--group-016.md#canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203002220302032-3131301330311110-0003002013001012-2003000113203113-3300113031010323-0222312211302001-2012120212310300-3233011032333123"></a>

## simple_service — simple_service / 303131233103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- simple_service

<a id="canonical-3033330332311101-3210113002001102-0301111101113100-3023101200113133-0000233331000233-1000221300201111-0133220110233111-3012103111100322"></a>

Type: `"object"`. single nested block, Optional.

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Upstream description:

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disabled",
    "enabled"),
  validators.ConflictingObjectAttributes("do_not_advertise",
    "simple_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"do_not_advertise\",\"simple_advertise\"]",
  "x-ves-oneof-field-persistence_choice": "[\"disabled\",\"enabled\"]"
}
```

Terraform syntax:

```terraform
simple_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001302321230311-1303111333132020-1320030010001303-0301102031323123-2100322301212331-1102300003031111-1102300323321122-0110213123021020"></a>

## Direct properties — simple_service / 303131233103 / 3

- [configuration](resources--workload--reference--group-016.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313): complete subsection reference.

- [container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213): complete subsection reference.

- [disabled](resources--workload--reference--group-017.md#canonical-1110023311313011-0202000212230320-0210021130232023-0323322102330130-3223121303212321-2033302130023021-3120001020331002-1031122330013020): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-017.md#canonical-0003300212123321-0030022111222023-1030022232221021-1132031212101132-3230032012113311-1303232022202331-3111103333310112-3310021021132230): complete subsection reference.

- [enabled](resources--workload--reference--group-017.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030): complete subsection reference.

<a id="canonical-1002131132203101-1131232101213330-0032322330310111-2211001301212333-2231300212010012-2332222020210322-3110330322123132-2002212121103302"></a>

<a id="canonical-1213301332310000-0002332303221032-1020131333331113-3001231100122001-3001332032130311-2130322331211032-0103022200130101-1211011032202022"></a>

## scale_to_zero property — simple_service / 303131233103 / 4

Type: `"bool"`. Optional.

Scale down replicas of the service to zero.

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

- [simple_advertise](resources--workload--reference--group-017.md#canonical-3310302312301012-0213323302330302-2012000032301123-3011230301130001-2021102001021321-2321111011331112-1000322300330133-3212013211230202): complete subsection reference.

<a id="canonical-1223223010302123-2302030103221200-2013310332031003-2212330232330232-3121323331032100-2022302101221030-3322220100023310-0221232333021210"></a>

## Next pages — simple_service / 303131233103 / 5

- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.disabled](resources--workload--reference--group-017.md#canonical-1110023311313011-0202000212230320-0210021130232023-0323322102330130-3223121303212321-2033302130023021-3120001020331002-1031122330013020)
- [simple_service.do_not_advertise](resources--workload--reference--group-017.md#canonical-0003300212123321-0030022111222023-1030022232221021-1132031212101132-3230032012113311-1303232022202331-3111103333310112-3310021021132230)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [simple_service.simple_advertise](resources--workload--reference--group-017.md#canonical-3310302312301012-0213323302330302-2012000032301123-3011230301130001-2021102001021321-2321111011331112-1000322300330133-3212013211230202)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003232031000132-1112101012021122-2111301233011100-2303300313332331-0021111001102310-1312010033112310-3030303101013113-3130111230021310"></a>

## simple_service.configuration — configuration / 132201300110 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.configuration

<a id="canonical-1232232331100102-2211301230213012-2322020311120000-3001332301233103-0320330130231120-3132333102232030-1333212310002333-3221332033223121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameters of the workload.

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
configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001013312312332-3202212322020013-0130330230001110-0022201300130121-2023310312003030-1312131132022203-0120022011021110-0110030031133113"></a>

## Direct properties — configuration / 132201300110 / 3

- [parameters](resources--workload--reference--group-016.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130): complete subsection reference.

<a id="canonical-3221032033001131-3320323132000030-1312031302201330-0213212231331200-0002133312022011-3333231223002030-2121332112133320-3230011213213033"></a>

## Next pages — configuration / 132201300110 / 4

- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121202202020123-0020123132313122-2122111101313102-0200221211223303-2112231013231131-1133200021111231-3112112220123011-1101333020131331"></a>

## simple_service.configuration.parameters — parameters / 301203002023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- simple_service.configuration.parameters

<a id="canonical-3023003010131101-3113213303201203-2310202100320233-3003231031023112-3030102021331012-3303003002220203-0301032312231012-2011223312103202"></a>

Type: `"object"`. list nested block, Optional.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("env_var",
    "file")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020131212333233-1321132321100120-0132132101122120-1322302220232323-0231103001221213-1133102120331231-3101000202012012-0032123312132123"></a>

## Direct properties — parameters / 301203002023 / 3

- [env_var](resources--workload--reference--group-016.md#canonical-0102122312313001-0221331223003110-2133110201320030-0012331100210230-0200113132332023-3121201230202020-1010310113323210-3103220100311122): complete subsection reference.

- [file](resources--workload--reference--group-016.md#canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102): complete subsection reference.

<a id="canonical-1223331203230310-3130231333210031-3211332202032223-0120222003222202-0310023011131020-3332001330301313-2233111120102213-1112003123101002"></a>

## Next pages — parameters / 301203002023 / 4

- [simple_service.configuration.parameters.env_var](resources--workload--reference--group-016.md#canonical-0102122312313001-0221331223003110-2133110201320030-0012331100210230-0200113132332023-3121201230202020-1010310113323210-3103220100311122)
- [simple_service.configuration.parameters.file](resources--workload--reference--group-016.md#canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0102122312313001-0221331223003110-2133110201320030-0012331100210230-0200113132332023-3121201230202020-1010310113323210-3103220100311122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313120121112031-3003113031110013-2320201301110200-1230302030220211-1331123100033223-0033213333332301-2023110322101101-2023330231103213"></a>

## simple_service.configuration.parameters.env_var — env_var / 133023210102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- simple_service.configuration.parameters.env_var

<a id="canonical-2332232003231032-0100033022223300-3212132001013332-0320210000212210-1311111231210310-2121300330232222-0303321222031001-0111331333313011"></a>

Type: `"object"`. single nested block, Optional.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

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
env_var {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323101332311102-3201230233311122-2020310220210122-3200223310010223-0023010222232013-1223033122121302-2033020023112111-2221030332011123"></a>

## Direct properties — env_var / 133023210102 / 3

<a id="canonical-0120033022232303-2012303111030203-0200102330312023-3122330232331221-0131112303100020-0322333001223223-2213221100100031-2103011310010211"></a>

<a id="canonical-0302011203300131-1211221003101300-3021210130012100-3210303030133022-2232013101111213-0103111311300203-1233032301200323-3003333233130130"></a>

## name property — env_var / 133023210102 / 4

Type: `"string"`. Optional.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3122002020102310-0313020230020302-2110032332020012-2023001302130333-3100213120310032-2120333322132232-2223023230232023-0300320122120122"></a>

<a id="canonical-1220333213131230-1223001323212322-1022113130113331-0131231031301031-0202220212312120-3202232122202230-1230000211303100-2133233103101130"></a>

## value property — env_var / 133023210102 / 5

Type: `"string"`. Optional.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1202103202312313-1220110002213031-2321201033130232-0113102120120300-1300112222232000-3012032113121000-0012130230313032-2002202021100113"></a>

## Next pages — env_var / 133023210102 / 6

- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000311332220331-0001223031310113-0312222111230222-1222311100311012-1310200103102013-2201330002102210-3102210311133130-3301211300222331"></a>

## simple_service.configuration.parameters.file — file / 301330133203 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- simple_service.configuration.parameters.file

<a id="canonical-1001032302332122-0310100122133230-0212100223222003-3010230130313121-1001033000113132-1231020332102030-2220111103133300-2210213013300300"></a>

Type: `"object"`. single nested block, Optional.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "volume_name")}
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
file {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023232121223022-1330003010200033-3321102311212101-2332322302030031-2323023323323211-2121111323301111-0132101021312333-1101323313022112"></a>

## Direct properties — file / 301330133203 / 3

<a id="canonical-3010210033130313-0312302212012231-2031002312110322-0320200011300130-1310233301223223-3133033103322301-1031033231001301-2003231320003121"></a>

<a id="canonical-0133000202230311-2013022120001313-0323300230123023-3321020231113221-3302322013112303-1012200021023310-3110020030212120-0021331132222122"></a>

## data property — file / 301330133203 / 4

Type: `"string"`. Optional.

Data. File data

Upstream description:

File data

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16384),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](resources--workload--reference--group-016.md#canonical-2101021012222331-3232010123232032-3303103311113010-2030322211032012-0103002110333102-2031001130221330-3231010210311002-0223201022223102): complete subsection reference.

<a id="canonical-0213311303231321-3332331213102132-1301001212001022-1331231121223220-1232323021110110-2213123102322022-1220301002303212-0202110322023323"></a>

<a id="canonical-3221303223312222-3112221110321111-0220310122232102-1211033321312022-1332202233211200-2310330020033331-2332111132301203-2010121232323213"></a>

## name property — file / 301330133203 / 5

Type: `"string"`. Optional.

Name. Name of the file.

Upstream description:

Name of the file.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1020102202211311-0022000202330231-3012011302303222-2000023221012032-1000113311001021-3003322001010120-1330222033303311-3010001300123210"></a>

<a id="canonical-1113101200323013-0123312321031333-0033101212132332-0033330331220031-1100231312001211-0303322132320232-1020100333322102-0333022310120003"></a>

## volume_name property — file / 301330133203 / 6

Type: `"string"`. Optional.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2021202212320322-2031320323120301-1220221022303013-1020112302030122-2303002203321312-0313131022303313-0030300222032220-3020121200232333"></a>

## Next pages — file / 301330133203 / 7

- [simple_service.configuration.parameters.file.mount](resources--workload--reference--group-016.md#canonical-2101021012222331-3232010123232032-3303103311113010-2030322211032012-0103002110333102-2031001130221330-3231010210311002-0223201022223102)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2101021012222331-3232010123232032-3303103311113010-2030322211032012-0103002110333102-2031001130221330-3231010210311002-0223201022223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130320330101330-2222133220210332-3320120333220033-1111300323222110-3322212013020132-0331312020121223-0200023113021032-1300231323203020"></a>

## simple_service.configuration.parameters.file.mount — mount / 203222303103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.configuration](resources--workload--reference--group-016.md#canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313)
- [simple_service.configuration.parameters](resources--workload--reference--group-016.md#canonical-2010201103131330-1012302010203000-0330231021220300-1301210002222232-1222023123330202-1322233233110310-2311103301211122-1202212302232130)
- [simple_service.configuration.parameters.file](resources--workload--reference--group-016.md#canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102)
- simple_service.configuration.parameters.file.mount

<a id="canonical-3212302022331323-2211221201222003-0123112130020313-0021121131303213-3002223111210012-3321032111020210-1020030120103103-1111213123200211"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301200100203221-2111312030020000-1020233021102130-0120230000113033-0220001221111211-2002131120022212-2020213002230331-3113010321312013"></a>

## Direct properties — mount / 203222303103 / 3

<a id="canonical-0003012212033002-0103213210031101-0321200331020101-3113210023133330-0302233113032010-2231203212002302-0012001312110210-3231032021002200"></a>

<a id="canonical-2313010302111030-2101010030021303-1031121133022220-0223321200133022-2032120232220313-3210321011133301-2213200323330313-3002333232123211"></a>

## mode property — mount / 203222303103 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3212330223330222-1333101133203312-1323121233110023-3230103101201201-1001213011133211-0030032332130231-3212133222021230-3323213112201131"></a>

<a id="canonical-2120311113233102-0313311033120021-3331213003231013-2333200201213301-1021210012123213-3021112323130231-3331330323113102-0120333232000110"></a>

## mount_path property — mount / 203222303103 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-0313113033320011-0122021333002211-0021301013220210-3112230121303232-3123030012020321-3203213133021130-0133000310031130-2023220213000122"></a>

<a id="canonical-0001202131233001-3002232210121003-2021210113333022-0202200112203201-2301321221222133-2323230120323120-1300302130332032-2022102132333323"></a>

## sub_path property — mount / 203222303103 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3120020332102101-3023200222023200-2230221003120030-0331130223212233-2132333132121201-3120310320031311-3333310303333330-0300213021210213"></a>

## Next pages — mount / 203222303103 / 7

- [simple_service.configuration.parameters.file](resources--workload--reference--group-016.md#canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101130232232331-2300013222111331-3013102321321013-1223111320232103-1233202033303202-0112101032220333-2230312332023311-2102021102212020"></a>

## simple_service.container — container / 233033011010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.container

<a id="canonical-0000332130000122-3333320032231313-0322210332113231-0003302103123222-1330101203111233-3131323011322223-2330000001200103-1031033200220213"></a>

Type: `"object"`. single nested block, Optional.

ContainerType configures the container information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("custom_flavor",
    "default_flavor"),
  validators.ConflictingObjectAttributes("custom_flavor",
    "flavor"),
  validators.ConflictingObjectAttributes("default_flavor",
    "flavor")}
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
  "x-ves-oneof-field-flavor_choice": "[\"custom_flavor\",\"default_flavor\",\"flavor\"]"
}
```

Terraform syntax:

```terraform
container {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212202301123302-1222020002301220-2223321211233103-0032213100300201-2123021232211211-3022320312002203-3030303330331333-2001223010333302"></a>

## Direct properties — container / 233033011010 / 3

<a id="canonical-1320122102200012-0000332123203102-2100330202122002-1331133230000220-1323231203132120-1113111032322020-3300332101320233-0110002002331332"></a>

<a id="canonical-3201102102112011-2131320300230310-1130102322010333-3202310323002030-1001310130312323-0103113002000202-1003003013021303-1121123022321303"></a>

## args property — container / 233033011010 / 4

Type: `["list", "string"]`. Optional.

Arguments to the entrypoint. Overrides the Docker image's CMD.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-1223111301003112-2002220311333111-3002231330301122-3211022012210103-1032210300312001-2322112323021120-1123010111130211-0220201310012133"></a>
