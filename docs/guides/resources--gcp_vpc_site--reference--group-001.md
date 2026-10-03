---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012302323220212-0133022003331301-2210023022211012-3020110333302233-0220322202120232-1020230021212210-2303012103231321-1202311101223210"></a>

## Property reference — Property reference / 020322303323 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- Property reference

<a id="canonical-2021021102112111-0032132133203300-2332210132020312-0300200202223112-3233031130311221-1023210201303022-2121323112130210-2022300311322323"></a>

## Direct properties — Property reference / 020322303323 / 3

<a id="canonical-0131020013203313-3030130102013100-0102323120231130-1101022012113122-2100300001122113-3020311111132212-2021220122122300-3321200123301113"></a>

<a id="canonical-2300030201200213-3133213312011112-3311111002200121-3311223131222303-3010211321230311-0210022303101130-2230103121210010-3022131011020230"></a>

## address property — Property reference / 020322303323 / 4

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-0120223131022023-3232101001310331-1103010012022230-1200131021301230-1312302111223330-3010233323332131-1111303112300132-1103200021202303): complete subsection reference.

<a id="canonical-1013210010031233-2200323310011033-3211210331312120-2201130033003111-3202331301111232-0202100010311021-0313103131022333-1201000321110131"></a>

<a id="canonical-2211001021002210-0210003323321012-3303313213111302-0023203302033132-1103012121102300-2130312300232333-1300012130100332-1010233133133332"></a>

## annotations property — Property reference / 020322303323 / 5

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
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

- [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-1103003322103331-3002300011231332-3102223020003212-3211211233001332-2032120010123310-0313231000122102-0303201323233323-3101332311030331): complete subsection reference.

- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032): complete subsection reference.

- [cloud_credentials](resources--gcp_vpc_site--reference--group-001.md#canonical-1312001102001300-1311230131021121-2213102022222331-2331123300121313-1323121213000112-0102121102020303-2231311102230221-1322021220133331): complete subsection reference.

- [coordinates](resources--gcp_vpc_site--reference--group-001.md#canonical-0132322102032221-0221031120301201-1222222320102200-3032133011203110-3012200020001320-2313113033220032-2222201130301012-2121000101223020): complete subsection reference.

- [custom_dns](resources--gcp_vpc_site--reference--group-001.md#canonical-3220300023200020-3013222211030303-0100203211331300-0030113331233123-3331132002222231-0000301012011231-2310300112031020-2032323332321030): complete subsection reference.

- [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-3310031330331201-2123120101221211-0113032011130223-0003312130311301-1303121010101221-3233302131102330-2311002110111331-1013112312302133): complete subsection reference.

<a id="canonical-1011110323011330-2220302212130333-0303030233122332-3331310112203103-1322032011113012-2233220323231011-0122133132013111-2011011001311003"></a>

<a id="canonical-1203121313132031-0012233012221032-0011130322331333-1022012311301130-2232310300023303-0313322233032033-0220132131330020-0031223230012303"></a>

## description property — Property reference / 020322303323 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0000111301212302-1203003212211201-1203211031010312-0331010131203022-0300333312121001-1022112031301131-0302131220301233-1123013210011322"></a>

<a id="canonical-0020333203212130-3001123220212332-2101201202011001-2132020202010330-0133121231211322-2102110100321231-3302133031210023-2213102021112201"></a>

## disable property — Property reference / 020322303323 / 7

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

- [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1131111322332132-3103020101000022-2211033321232233-2030233213233230-0131222130003033-0220320200112322-2313123101130311-0133321200322303): complete subsection reference.

<a id="canonical-0121113310010011-3200323201200333-3212211231012002-1230210020301313-3120322130212310-3323023101212223-1313001311012130-3210313220020122"></a>

<a id="canonical-2001021323301331-0312013101120221-0133320123010010-3202320300012201-2020213323323112-3220000331203202-3311133102003122-1032101230122232"></a>

## disk_size property — Property reference / 020322303323 / 8

Type: `"number"`. Optional, Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1101221123132310-1010130003200313-0032320202321131-0203333233313013-2011212112020022-2213003211113113-1032113120210301-2132310133203302): complete subsection reference.

<a id="canonical-2022012110322200-3231220032032201-2000131013001030-3330032133213233-3310302312330021-1113223200320310-1100101231212033-0101232011201233"></a>

<a id="canonical-2133021303001023-1003100201032003-1100220303320221-0333333300003313-3133302233011031-2312300210321123-1212032030230101-3113030010203200"></a>

## gcp_labels property — Property reference / 020322303323 / 9

Type: `["map", "string"]`. Optional.

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

Upstream description:

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{
  validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":40},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":127,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"127\",\"ves.io.schema.rules.map.max_pairs\":\"40\",\"ves.io.schema.rules.map.values.string.max_len\":\"255\"},\"values\":{\"maxLength\":255,\"type\":\"string\"}}"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-1023010333222331-3121302012023130-3321223333203300-0202002133100133-0121021101032031-3222032310110300-0223113201331121-1033201223101032"></a>

<a id="canonical-3200212220101311-3123112120112110-0200203020320022-0310031013220012-3131111221212212-2203122220321102-2110232101211303-1322001300003032"></a>

## gcp_region property — Property reference / 020322303323 / 10

Type: `"string"`. Required.

GCP Region. Name for GCP Region.

Upstream description:

Name for GCP Region.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0300321131132100-1323012111132200-1110332022200231-3123113101123320-2113011030110102-1022120202011201-0131323321303022-1303223333202013"></a>

<a id="canonical-3011301202120012-2311011320013000-2211203111100330-0011020003033222-1230321313001313-3031013330200300-3323031013220003-3211222202332311"></a>

## ID property — Property reference / 020322303323 / 11

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100): complete subsection reference.

- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-2221222220302122-3232110120322221-3100201231013000-1213023302033003-1313222101212322-2133001213332231-0312122202112110-2121222013023030): complete subsection reference.

<a id="canonical-1122210310200102-2321323213130113-1032300030312331-1012022323131121-1012102233322210-3002311313033203-0231013301102331-2301023020133020"></a>

<a id="canonical-0010331012320310-2020130301121020-3130202011220231-1210323222303333-3330313313212201-1300303323130311-2303130112100121-2032033012301011"></a>

## instance_type property — Property reference / 020322303323 / 12

Type: `"string"`. Required.

Select Instance size based on performance needed.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-2101212102330031-0010112330122012-1102313232223312-0222112303201021-3012113311101013-3130313120212310-1030123002232113-0300133312323323): complete subsection reference.

<a id="canonical-2101102011223332-0323021220110131-0222123103022111-1003113313133332-0310332303232331-2210000012120013-1021330100101300-2103021222001310"></a>

<a id="canonical-0323201123323130-2012312003232302-0332001222132222-3211031233111120-0033122312302332-0231132233122100-1000313300331023-0133033003233100"></a>

## labels property — Property reference / 020322303323 / 13

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
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

- [log_receiver](resources--gcp_vpc_site--reference--group-003.md#canonical-1023013120132311-1101001323331302-1302211232003313-1023011331032223-1301113322212223-1030321303032302-2211010200203113-1223301000320112): complete subsection reference.

- [logs_streaming_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3001101023203132-2313313210112200-0002300333010120-1210100002212200-1202021011233220-1222101232112112-2003010130010021-0013120113331012): complete subsection reference.

<a id="canonical-0322223203122102-1330212020100230-3212221100012111-2301110122322232-0130332002032012-1331301323223011-0230232130003222-3323002213300333"></a>

<a id="canonical-1113122111122310-1120023323031332-3020022311303113-1133202100232302-2213001303211102-1221113212000032-0212303212023202-1331111230332000"></a>

## name property — Property reference / 020322303323 / 14

Type: `"string"`. Required.

Name of the GCP VPC Site. Must be unique within the namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2002220012311032-2210330133321031-3310211030011211-3031013023203333-2302121121230031-2221032132101030-3313011130212201-1322101232301122"></a>

<a id="canonical-0121322030013201-3230132101230300-3323330132322002-2313022010011331-3032223032302333-0121013320001332-0230121122112120-2331331013030131"></a>

## namespace property — Property reference / 020322303323 / 15

Type: `"string"`. Required.

Namespace where the GCP VPC Site is created.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-0302011120033022-2021220221211032-3112222301002021-1021012200222002-0322233013032322-0132210231022013-1300001101112021-3332323033020031): complete subsection reference.

- [os](resources--gcp_vpc_site--reference--group-003.md#canonical-3222223210301122-0312211222330230-3122221002123233-0113101130301311-3320303012132320-0302323231130332-0123201033332330-2211021130022023): complete subsection reference.

- [private_connect_disabled](resources--gcp_vpc_site--reference--group-004.md#canonical-1200311111133100-0020033221223212-3021002202113332-2301121021212132-0232212133303310-2231010113102221-0033221003131012-2022210321113201): complete subsection reference.

- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133): complete subsection reference.

<a id="canonical-2221333211121223-3100331330030210-1120033101130210-2003113111111330-3023332233122030-3200210331311003-2121310233223101-0032123133313012"></a>

<a id="canonical-1013001301013212-2101123300100302-0000030312302201-1330222300002230-0223022131332321-1310132300331011-3102212323032033-0023301211001110"></a>

## ssh_key property — Property reference / 020322303323 / 16

Type: `"string"`. Required.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](resources--gcp_vpc_site--reference--group-004.md#canonical-0212133321121330-1301131103023331-2123132213311103-0031232100002231-1233033231110220-1102201033230030-3020012013122301-2110012031323002): complete subsection reference.

- [timeouts](resources--gcp_vpc_site--reference--group-004.md#canonical-0202212023111322-0302222202030003-3121101231202200-2213031121020101-0121233300121223-1032213023303312-0220022023213022-0231230022223121): complete subsection reference.

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113): complete subsection reference.

- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-2220233032031010-1323100300312032-0222201202010130-3130203022232222-0300133130321233-1132212212212321-0322012021310311-1132132110223100): complete subsection reference.

<a id="canonical-3302131113102030-0200022232113110-2111222020122003-2211312103202212-1003220231233222-3303100322333111-0213022001333302-0011023202202311"></a>

## All schema paths — Property reference / 020322303323 / 17

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--gcp_vpc_site--reference--group-001.md#canonical-0131020013203313-3030130102013100-0102323120231130-1101022012113122-2100300001122113-3020311111132212-2021220122122300-3321200123301113) |
| `admin_password` | [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-3003220023201110-0123012222203123-1123232121100300-3130110132221221-0003313012032303-3232221111101211-1120011000101222-1122023112003130) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-1332023311103202-2230103223031121-2121233332332103-1223132013113101-3311120330022011-1000121101313012-2211023022132033-1303032222130010) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](resources--gcp_vpc_site--reference--group-001.md#canonical-1312023222220223-1012130203310030-2013313201332031-3003013233310030-2230110002211332-3201003311300111-2002001221021300-1111132213211100) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](resources--gcp_vpc_site--reference--group-001.md#canonical-0312331110211323-2320303321232213-0120201001023022-3002220133003030-2132221133130002-2101110121021311-1232210312301322-0203023001222333) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](resources--gcp_vpc_site--reference--group-001.md#canonical-3233031121121033-3213233102322320-2203132020211022-2133020100311321-3331331213120233-1010111200211113-3230000330030120-0100123031032012) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-0023012333203103-3202010223113032-1312333031222113-0030203120110101-1222113222031302-2131202332303222-3311101331332030-1023012321301102) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](resources--gcp_vpc_site--reference--group-001.md#canonical-3113320223122002-1132200102120012-1022212120000302-3030022130123131-2101231303213110-2103332021313330-3223311112212332-3201110330232133) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](resources--gcp_vpc_site--reference--group-001.md#canonical-0323200001031310-2121103020112112-2201003200032121-0002002310032001-3221003112230320-3103312200121130-1302311022133022-0210123001020002) |
| `annotations` | [annotations](resources--gcp_vpc_site--reference--group-001.md#canonical-1013210010031233-2200323310011033-3211210331312120-2201130033003111-3202331301111232-0202100010311021-0313103131022333-1201000321110131) |
| `block_all_services` | [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-0010203330220110-1321221020003123-0313333201303011-3320312311130233-3333101130013200-0111022232312012-1302103101110030-2223213332030112) |
| `blocked_services` | [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-1000323100231311-2301320123232312-2002123332011201-0233121312022123-1011032320123203-1103333123203103-3003223032211302-3313232121221332) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3032113103131331-2020313011111333-1023323013120331-2300012103020012-1113012330021330-0132321102332203-1300012122032222-3123330200221010) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--gcp_vpc_site--reference--group-001.md#canonical-1112230303310220-0031101202233200-0201131111111021-2030022133021000-0223013233131031-3321202203222110-3002033332230231-1121031122113330) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--gcp_vpc_site--reference--group-001.md#canonical-0331102313120332-2313220333130202-3230313311321311-0012332232321210-2321231130001212-2301320301310032-3222032122202002-2231200031101020) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--gcp_vpc_site--reference--group-001.md#canonical-3103133302320113-2131131231122123-0010102212113001-2032203222102010-2302030112321200-2012213113131011-2020330211132100-1213330332013002) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--gcp_vpc_site--reference--group-001.md#canonical-0212221312333121-0120321230023130-0121033203122102-1331231121320122-2022001203203130-2011333233202333-3332002102203122-0230323122203320) |
| `cloud_credentials` | [cloud_credentials](resources--gcp_vpc_site--reference--group-001.md#canonical-2212230332101332-3103012112233333-1203033220233232-0101223001101320-2210013131021123-2003223030210221-0133303033131021-0201320211213022) |
| `cloud_credentials.name` | [cloud_credentials.name](resources--gcp_vpc_site--reference--group-001.md#canonical-0210302331122110-0323000100132212-1300033130221332-1221003213212321-0033233001232333-1001200121233333-1323011133310231-0111210333033111) |
| `cloud_credentials.namespace` | [cloud_credentials.namespace](resources--gcp_vpc_site--reference--group-001.md#canonical-1221300333301322-1320131302001102-3003103211100111-3211112231312110-3103101301331220-0222122330033331-1033233022312012-3203223213101011) |
| `cloud_credentials.tenant` | [cloud_credentials.tenant](resources--gcp_vpc_site--reference--group-001.md#canonical-3101212133320031-0122203131232031-1211013120220033-1033231211332121-0321110203300322-1203121302023313-1230133301301031-3103022200331123) |
| `coordinates` | [coordinates](resources--gcp_vpc_site--reference--group-001.md#canonical-1011322230220003-0302003212122331-3012223120323233-3131211303002303-0310020101322313-1023022223220102-3230322000001132-2203021002012121) |
| `coordinates.latitude` | [coordinates.latitude](resources--gcp_vpc_site--reference--group-001.md#canonical-0010200321313230-3131211010012101-1322003022301201-0111200320200221-3312313331312122-1310122102301002-1000221301102100-3221302000012100) |
| `coordinates.longitude` | [coordinates.longitude](resources--gcp_vpc_site--reference--group-001.md#canonical-3123320120201112-3022203313103232-1220112011030001-1031323222221032-2301330212223022-3210100002231312-1233303132003111-0032223232322302) |
| `custom_dns` | [custom_dns](resources--gcp_vpc_site--reference--group-001.md#canonical-3002030232232100-2033020121311131-2133200130002313-1211330311002133-0010020103231120-0101101131302212-3001202113331320-2122033003121233) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--gcp_vpc_site--reference--group-001.md#canonical-2232000032233212-2323112013020013-2201111330201213-0223011322133203-2300303233013122-3131033021212112-1123120123032130-3321321331013133) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--gcp_vpc_site--reference--group-001.md#canonical-2232030010333330-0012212102213311-0221200320033332-1220100022201320-3033331102012320-0121233011020100-1133321333233110-0201321021001030) |
| `default_blocked_services` | [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-3221320222002002-3132322013032031-1231302031321113-3320022220113332-3202301310311303-1213132233212112-2000011221113330-1331212333223123) |
| `description` | [description](resources--gcp_vpc_site--reference--group-001.md#canonical-1011110323011330-2220302212130333-0303030233122332-3331310112203103-1322032011113012-2233220323231011-0122133132013111-2011011001311003) |
| `disable` | [disable](resources--gcp_vpc_site--reference--group-001.md#canonical-0000111301212302-1203003212211201-1203211031010312-0331010131203022-0300333312121001-1022112031301131-0302131220301233-1123013210011322) |
| `disable_encryption` | [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1211200031130233-1323333201331301-0322231221101222-1120210211030120-0120010001323132-1212132221021213-0323103002020013-1303120013230211) |
| `disk_size` | [disk_size](resources--gcp_vpc_site--reference--group-001.md#canonical-0121113310010011-3200323201200333-3212211231012002-1230210020301313-3120322130212310-3323023101212223-1313001311012130-3210313220020122) |
| `enable_encryption` | [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1201103122112231-2220103102133220-0020113032201213-2220030020301213-0121101300223203-1022230101003312-3002300313221133-2301021211022321) |
| `enable_encryption.kms_key_resource_id` | [enable_encryption.kms_key_resource_id](resources--gcp_vpc_site--reference--group-001.md#canonical-1331133121310212-1132003312013130-0230222310313103-2121122323221302-3000132331102322-3122022131002011-2033001012030110-3112020302202110) |
| `enable_encryption.kms_key_ring_id` | [enable_encryption.kms_key_ring_id](resources--gcp_vpc_site--reference--group-001.md#canonical-2333220201230320-1023313303133030-0031220230301211-3111111003033122-0110213200323301-3303202002311232-1313200023322233-2023211202130000) |
| `gcp_labels` | [gcp_labels](resources--gcp_vpc_site--reference--group-001.md#canonical-2022012110322200-3231220032032201-2000131013001030-3330032133213233-3310302312330021-1113223200320310-1100101231212033-0101232011201233) |
| `gcp_region` | [gcp_region](resources--gcp_vpc_site--reference--group-001.md#canonical-1023010333222331-3121302012023130-3321223333203300-0202002133100133-0121021101032031-3222032310110300-0223113201331121-1033201223101032) |
| `id` | [ID](resources--gcp_vpc_site--reference--group-001.md#canonical-0300321131132100-1323012111132200-1110332022200231-3123113101123320-2113011030110102-1022120202011201-0131323321303022-1303223333202013) |
| `ingress_egress_gw` | [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3012232203320333-1330203033310313-3102123222033211-1333321322230123-1322023301131002-1212320311301032-0121011000200030-3013011002010022) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-2333230122003302-3020023031221130-3122001123331011-2100011230312313-0223332121311211-3122201002010201-3312021123221012-2222313023210112) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0113111111031113-3011300103132031-0132203020202210-3302311221232030-3033120002030330-1020320223322022-1131311002022121-3310223231101212) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0201032031003300-1221302300302013-2232300003231010-1233022122220330-0221020013303231-1010321303301100-1111022330112030-0221010022010011) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-1012123300312120-2203100130103212-2203120221010331-3110111320031002-1312231210023212-2003231202211203-3211212131212203-1012311300203313) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-1103020120020333-0133011132210120-2302023101322013-0031232111133002-3333320130303111-2030233003312203-2213211233233203-0302221200121003) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0332012032001200-3323002331020200-3302001010020303-2333021313131100-2002223302322031-2232112013300310-2102031102311101-0222300222211230) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-3200013202113113-2230122330013220-1030201130320103-1103020032213200-1203011332201122-0301332103330332-2202320210013220-2302332003031232) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0331113031323311-0302230323121120-2213023100102233-0011331333112332-1232311320303130-3231220331211232-2321100213312321-1222102122232302) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-3232103212300032-0221100302111201-2301210223120100-0022012022000102-1202310030031133-1303131012031213-0301202100110011-3230123000100300) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-1021102023131202-2003121213332013-0203031230211021-3320221303330120-1020233221320003-3123002101002233-2202020023130201-2032121003132333) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-2220200022201113-3011311120102223-0130202031103110-1033321231123130-2013113113011310-1131322021003020-3332222311223131-1011000203002013) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-3211203032233303-1312101213323132-3103000223332010-0203121032311103-1133022121302033-0103203333123203-0033303123012302-1221203121113111) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0111022122330223-0130012111031011-0130201123231132-3302120203110221-0333230102232203-2303123121123122-0200033232330002-2232000320223121) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-0122111121301330-2322033222320031-1021233101222112-1231133333110332-3012110132011100-0333212023212312-3221003232133201-1021013011113113) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-0202300221000013-0001102023002211-0333110203002123-3003331330200021-1333032022212101-1212210213212330-2233111021101111-2122210233303103) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-1100013301000323-0020330133001213-2320323001331120-0011221310033133-0111022021112033-2310022213130132-0021313213313132-2023113021333303) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0000101232202101-0033232022203230-2333133131031103-1033333222310212-1223121201001300-2131031222231130-1312311310320320-1311221033030101) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-0013111113313320-3310111221122302-2202023221202313-2302332301022330-0013330030302110-0013031303110330-1222223313020122-0033213302013022) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-1220102300233212-0033102322103310-3110200212033121-1230001002231031-1112233232113020-2300302010202132-3102021322010220-3223121231021131) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-0102302312122200-0130122133223222-0301013031112303-1022122210320131-1311212023102213-2101332101103000-2221213310011022-1133103200213333) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0131013121320031-1021330003231201-1321310010323212-2300212032010030-1202213001032133-0321200323001211-2132022122113302-2113333321012033) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-1010002130330031-2232333321110213-3002332001001010-1331203303003311-3310031321312112-0010010131232030-0020110100113321-1232032013101202) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-2130012201132203-1323110111213010-3212223302030333-0332220231202332-0021331212233330-0122111021203333-0131110331231320-2301222222033230) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-002.md#canonical-1221312311231120-1300332032323121-0111021030023320-3013313102203111-1020231100321023-0210103213111301-0201113031230001-0321101131231311) |
| `ingress_egress_gw.gcp_certified_hw` | [ingress_egress_gw.gcp_certified_hw](resources--gcp_vpc_site--reference--group-002.md#canonical-3132033023002102-3210110010221302-1333203113210213-0220330112001232-0211312221003312-0210312312300022-1332313031203300-2131312302212000) |
| `ingress_egress_gw.gcp_zone_names` | [ingress_egress_gw.gcp_zone_names](resources--gcp_vpc_site--reference--group-002.md#canonical-0133302222332222-2202131313133232-0021101120202220-2212120212111202-0002230322310133-0221122010301101-1320031321130032-1012101110332301) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3312020302221223-3233230233221013-1003302301021032-2102331210333330-3000322301230103-3111033220210302-0213333213102021-2012002203000303) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-2202111301201000-0231131220223032-2332033001022031-1111112303303220-1331023030303003-0022010330032011-1110112300203011-1213120111001321) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-1233223030331133-1100230033233331-1001233013013133-1003012020021032-2122130103312313-1010301213100133-3120220010232101-1112231033302002) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-0300221312010132-1322212103213130-3221200232130010-3300021122310331-1221211231102011-1222021020101322-1210113033212310-3200310223310101) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-3311120220212112-0010221201133212-0000120210221010-1123302303223322-3202012032332231-2223203011110200-2310320130231032-3000122101212201) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-3202301133200032-2201203030202231-1323012233113301-1311120121030032-1132321031312032-2022201311202033-0122011013333101-0020303001303331) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-0103323230133113-2132332210202303-1022022133031312-1030222210312002-3033223022122003-0301202102231211-1023110303122313-2310320023131113) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-2111110123220130-2110322123032033-0223001310130311-0310222031331011-0120130032231032-1333123202232231-1100231313020300-1031313211130303) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-2323303322313122-3220012300033111-3131301210220300-2121331212111023-3331013303101023-2112221031200010-1121213130031122-2202311230310331) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0302011021031323-1221202111210210-1011230303230023-2100223332111300-0032303131032232-2032130011212321-2110213301202330-2300312023303310) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-2230032020231310-3313320233113112-0023123030323222-1222213111202122-1123231322020023-2131011002233322-2323022313212210-3320331123310111) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-3121332013013313-1231213322230321-1001100031220322-2311211311310030-3222002332021302-1010320331200133-2030022203333112-3332211120312012) |
| `ingress_egress_gw.inside_network` | [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-1220201310013012-2332112312123201-1011202121130233-1121133233210021-1123221333300131-0011123322033122-1033122022013233-3332302113303311) |
| `ingress_egress_gw.inside_network.existing_network` | [ingress_egress_gw.inside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-0203312210210220-1000023000210202-2032310331331032-0223233223221021-2232110121210222-1113101123210003-3112010313122223-3030333133022111) |
| `ingress_egress_gw.inside_network.existing_network.name` | [ingress_egress_gw.inside_network.existing_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-1102021222030201-2030112310300131-3212121101223110-1302221003000022-3302023010322131-3001002020110100-3312331323230033-2121031003333120) |
| `ingress_egress_gw.inside_network.new_network` | [ingress_egress_gw.inside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2012233322012120-1333113001332212-0223203230123313-1110122223023301-2103123213332021-2120221231112332-1322323100131322-0021202212122021) |
| `ingress_egress_gw.inside_network.new_network.name` | [ingress_egress_gw.inside_network.new_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0203212311021121-0203011233003001-0032001100200010-0301302000321122-1032202130030302-1113320233033331-3201000230120000-2203311203022210) |
| `ingress_egress_gw.inside_network.new_network_autogenerate` | [ingress_egress_gw.inside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-3130011201120212-3020230122023232-3222133003302012-3110231310322201-2100130230313123-1323231103022003-2213021332102100-2122202303123321) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0133112210023230-3021200312011212-3313310011222011-2013120121231213-0103202113122323-3332103012111213-2313220033121303-2011020002310033) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3301310210223300-3213012303212002-3012220032223330-2123112133313103-0012000330012030-2101212303221100-0313131113012321-3020020001301102) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3200112120011022-1100223203333232-1301030102233010-0000332021011002-0321010111111302-1312030333132021-3322320000222103-3003002032000213) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](resources--gcp_vpc_site--reference--group-002.md#canonical-0001221013223201-0130023310210122-2210033111201221-1102302231311323-1123122102010233-3011132320310200-1321221231020012-2020022302112333) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-002.md#canonical-3220211310333232-0011032101322303-3132022210100201-3110003020301013-1112031001220211-1023130223111302-2111020322202301-0030022231222133) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3332131301221222-3012201130130202-1133103213313201-0312321303200123-0131220230111001-0120312001101322-2313102011330023-3111313313312031) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-002.md#canonical-0111101213002322-3133310002003222-0010033001100120-3123111231203210-2233030003030131-0331331131013122-3103301122123120-2023122233230121) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--gcp_vpc_site--reference--group-002.md#canonical-3013210303222031-3120232331330302-3221123313132220-0332300033133000-2022333231211320-1220301202211311-2231310232112130-2323332121223103) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0331312001212012-2310003003123021-2333113023103203-2132012132101201-2022101321122021-2200210003313033-0220122222222213-1210110123130322) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-1222333002213221-3000202203331202-2322330222011323-3213233001221333-0231000101120320-3222200211333021-1212211020110032-1132221030121112) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-2221030031221101-3010230313310201-0303300201332213-3223012002233331-0133031212010201-2033111201100103-0231000300322201-3310110020101111) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--gcp_vpc_site--reference--group-002.md#canonical-1230200212302333-3100323231300321-3032010200133010-3022013132210201-3311131233130023-1232103313301222-1012222210210113-0231002130123120) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-2113123313021012-0022131113003100-1213131122123210-1302330231303231-1132020013210232-2000022222130322-0202122030221023-3203332120013221) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-3020013230200310-3330320213132302-0321100103222221-0000202103110232-1220023310023310-2021233002321310-3332003310101301-3303212101131132) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-1223001332001211-0011011220211030-3000032011232001-2101322031211201-1330233031133312-0201201132031111-0000133301212133-1232113003301232) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-0111233323303322-3130131323001210-3233203031013220-2021033202201031-0023313130300101-0203021132100113-1102122133223130-3032102221231012) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-3001120313121113-0130311321323001-3231110210010020-3002203100132233-1221231133120332-1233111320312210-1102132332003301-2323322212330313) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-1310202121130231-2311023200023233-1203303203013223-2300121233012022-1012202333012120-1120211113203132-2332123122333333-2312203203101321) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-1020202310020110-1231011321201001-2012022111303213-2031120210120033-3102212022003122-1000133123210300-0321103012220200-1323011102022013) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-0112030332032011-1212020311011303-0213331221220201-3013110303120322-3230310103101110-0103133102011301-2110333303010313-1002100311310321) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-1032122123033232-1020000101211023-1222030100201232-1233312002330133-2031230001230232-1010120020013220-3323031201211130-2020322130123021) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-0212311323003120-2222121111100212-1012321321322111-2131123321321021-1322211012200031-2123033130221211-3000111111200020-1221233123130301) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--gcp_vpc_site--reference--group-002.md#canonical-3022302103133000-0000101130203310-0013303120123331-2003203130330303-0313101313333203-1210001102303030-0323222202010231-1013132210132020) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-3311220320031023-2130210220123312-0032220220303332-2213033122201123-2123231030013210-2230202002212302-3130320032213301-3331231101000320) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-0211110122301133-2230302231120112-2330212022221020-3123212110311220-3321031113202030-2113202333131213-3310210201000303-3223212320122032) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--gcp_vpc_site--reference--group-002.md#canonical-3233013000222222-1202112220322112-2110100120132321-0213301222013233-2121310320332201-1213113023331030-2020222210201132-2222321332012213) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--gcp_vpc_site--reference--group-002.md#canonical-0230011031302032-2311211013102332-1011230020331010-2011102101230023-2302221330221031-1120300202100133-1210130212103200-3111022330332013) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-0020110003112021-0022202020320313-1202323312300201-1103212222102332-2203012030113100-3113320002113020-0211102012301023-1102220302212130) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--gcp_vpc_site--reference--group-002.md#canonical-3010103222202103-0221220033131331-3003302013000021-2200101310130202-2311023100101021-0133023202302122-3022131313220200-0303123201023212) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--gcp_vpc_site--reference--group-002.md#canonical-1100112102210001-0333300200031321-0211301330131232-1130003301021320-2022331020313000-0310201112201012-3112202000223021-2310113330321200) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-2012312322322001-0302332013100301-3233111022333101-0313032011313103-2012301130300033-0013122132213333-0122321123212023-2130320202002310) |
| `ingress_egress_gw.inside_subnet` | [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-1002002222303002-1033332022232012-1132323323301123-3000001323010122-0330000310322033-3013111311001333-1311302322001003-3031333233203121) |
| `ingress_egress_gw.inside_subnet.existing_subnet` | [ingress_egress_gw.inside_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-2313032121200102-0230111103200121-2312010320123330-1332222323102211-2231302022321331-1032231113211022-0121023311000233-3011030121110012) |
| `ingress_egress_gw.inside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-002.md#canonical-0101020313230300-0102321212111323-0111300120202132-2123302021123320-1122212032231130-3311201120121233-1300302211000213-3303222223302123) |
| `ingress_egress_gw.inside_subnet.new_subnet` | [ingress_egress_gw.inside_subnet.new_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-1230013030223311-3311113110002232-2203032300322310-2220211133121032-0231010212102122-0013131102332302-1000113110222133-3123032012011110) |
| `ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-3333130303110311-3320133302131120-3303322123120300-3101300102311000-3222323320331212-2331313121100333-0121311133211003-1012210013202103) |
| `ingress_egress_gw.inside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-002.md#canonical-3332111002023131-3001200120002310-1333032301100101-2120333020333200-1233011300323221-1320122113123311-2133303021331031-3211332110302322) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-002.md#canonical-1110303003300303-2300033003302303-2220201231330120-2321123020302000-3333122221010013-1310112220332313-2223112201010233-1220121332330230) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](resources--gcp_vpc_site--reference--group-002.md#canonical-3102003113223311-0133133332230333-0103321130202321-3302301010332120-3110021112012230-2333312213013203-0203000302232033-0012330001200023) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3012032331120131-0123300301311331-3333110130031032-0033010211103203-2203331101221222-0100020013031302-3012323101120232-2220002303311233) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0001210102231111-1222103010132311-2122232301323332-2213132300032113-0303121201210320-2113032110003130-0301033303210030-1102003220032232) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](resources--gcp_vpc_site--reference--group-002.md#canonical-0320101233130300-3320331110011121-3212023231133032-2012311203220301-0010221320330301-2101301010222330-0120320132102021-1321121122012000) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-2113001200133131-1112322132232023-0120020231113001-2112133003331000-0202021010222221-3011130233330310-3032002132001222-2103121110030230) |
| `ingress_egress_gw.node_number` | [ingress_egress_gw.node_number](resources--gcp_vpc_site--reference--group-002.md#canonical-2122103103322113-1112020030111233-2333313300201131-0101133232122230-0322300033111101-2221322232023201-3021231232013113-3223222031212212) |
| `ingress_egress_gw.outside_network` | [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3203312133220231-3022221123221333-2023111002132232-2032022120111022-2210223002301200-0023210201022223-3112201223113331-1133020302101221) |
| `ingress_egress_gw.outside_network.existing_network` | [ingress_egress_gw.outside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-0201313200130011-0323213333231121-0231110013103313-2230012232221310-1102021312112320-2010201302032203-2001100122333111-3132121322122202) |
| `ingress_egress_gw.outside_network.existing_network.name` | [ingress_egress_gw.outside_network.existing_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0202321222332323-3012233031030221-3121233132223102-3320100001003200-0231302312003132-1110130312121330-1101122211222210-3220310321331200) |
| `ingress_egress_gw.outside_network.new_network` | [ingress_egress_gw.outside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2201123223203032-3301012133102120-3322100331002123-2300313110322023-3031001230031103-2311333221012222-2320112000222330-0330032030031120) |
| `ingress_egress_gw.outside_network.new_network.name` | [ingress_egress_gw.outside_network.new_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-3031120022012023-2121133022110023-1031110003320132-3011033202132000-2130000223023310-1203102333321302-1012023321002201-1203312332303020) |
| `ingress_egress_gw.outside_network.new_network_autogenerate` | [ingress_egress_gw.outside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-003.md#canonical-1310111201210310-2001323110131331-1332131213222122-2312221031203203-1312303010122221-1322122310110003-0100211220023231-2032033013310301) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-003.md#canonical-1122322302300120-2233123133203020-0201323330201001-2123020231203301-1300021120110222-2223012220221222-2101133223232102-2020332020100001) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-003.md#canonical-2321022220201132-2100020133312201-3310103320022003-1201310033120131-0030031013321123-2013331222111100-0232120110031001-3223223331221030) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-003.md#canonical-3222231203223010-0211103222133223-2103102332021113-2313311122300212-2302101311200212-0111112000233201-1300030332331011-3023110233102310) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](resources--gcp_vpc_site--reference--group-003.md#canonical-2223131130133001-2213311212002033-1333232123230100-0031001332123102-3320311211231033-3000032132301011-3230221321230111-2113233111311221) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-003.md#canonical-0222233110033020-3220000023110002-1312032122122323-0310123031110300-0102023202200321-0220033031201322-0202011021012333-1330203021312213) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-1313213003203200-0303111231201102-2321030230030123-0122012122320201-0000203332212330-1313320133032303-2032233322321323-1121132131301020) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-003.md#canonical-0302022223302022-3200123230210001-0321012012231013-3002213332110303-2203233333233221-2112122112112131-1210110000310330-2032123130010213) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--gcp_vpc_site--reference--group-003.md#canonical-2212223231313120-1031130310322213-0133233310132021-2133032320300331-3312222103212230-1222120312012112-3303230221230202-2213212111300310) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--gcp_vpc_site--reference--group-003.md#canonical-2313022212130201-2031323313301011-2211310200002212-2101211233302313-1323133320000321-2121222220333300-2130312310233120-1203233222123232) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--gcp_vpc_site--reference--group-003.md#canonical-0200103233233231-0331123200012210-0210202311303221-0333013102310010-0121232320120122-3130001321312330-3031113311223203-2323331223332020) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--gcp_vpc_site--reference--group-003.md#canonical-3130103230320331-2330002030100132-3331232132110001-2333032203200212-1033111202110013-3102012201123113-0323311222010300-0201120200003230) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--gcp_vpc_site--reference--group-003.md#canonical-1331311232332031-3113201300110332-0022232031013032-0022001303030031-2330330121003101-1111203331011121-3200223200132303-0311113200001102) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2003121002322102-2103333333001301-0201013030322302-1112301002213230-1100301022233112-3223030001000003-1023022000133123-0003211130200133) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-2132302022010310-2222200313021111-2222331022011011-0111000110212000-1010000330131010-2030310011313331-2220200100110001-1121212222022020) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-3010132113101010-1100231112133331-1130210121000331-0233110113310012-0201110211220032-0020111220103121-2131031102130031-1022201131301131) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-1211203311300021-3301322031121332-3110322322323313-0122113001123322-1003211120101013-2132120200231222-2122333103031033-1331211010200212) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-1120120222130330-0310231003311032-2310111131222033-0222233210032011-0100310332023201-0001320033330103-1100100113130322-2121210102203322) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-1010310202213320-1221010002133022-2013100001233201-3031110121100300-2221031201022233-2323210220233012-0210023023123002-0030100123000313) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-1302312111113321-3220000223201233-3302223201111232-3120302300221003-2110100000210113-0021003213113030-2331231100331301-3302322233221333) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-2002231023233321-3103211120001312-2210031120210210-3323223022310312-0203112121333301-0302112122233210-0200101300000021-1232123133230322) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-3013131300223231-0321200113021321-2200032321003313-3001222111211220-2100110200233312-3111300123101123-2102021113300321-0020032000030212) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-1102102020022213-3203101331030123-0001021033301201-1033133101103132-0321233300230101-3111001022131221-1332103123320212-2021002012122022) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--gcp_vpc_site--reference--group-003.md#canonical-1210213122000233-0003330002010202-2121212120201323-0001003133300113-1021011030331112-0033120130121223-0200312020130033-2200112121221010) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-2101202312031301-2300232123012123-0333333100232311-3122313132310023-2012011230112011-2003301012122033-3102321323122221-2020110300021020) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-2103333032023223-3232000030130112-0022130121123022-3022303321133001-3221321002232032-3131110110122322-2103321031230321-0331201033010323) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--gcp_vpc_site--reference--group-003.md#canonical-1030212212321001-1132122012303231-2000011111100333-1113012310112132-2313222020002201-2310223030023001-0310102231200031-1332111123320112) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--gcp_vpc_site--reference--group-003.md#canonical-2133231120030132-0033033331122110-0020211113123221-0331010211013321-0000110120221301-2103023030320122-2101031322122033-3212130011023111) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-0211310221113100-3232311210002312-3322221031103023-2332221030300233-3230323312313030-3132310023310310-3303331202302203-1101031122023330) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--gcp_vpc_site--reference--group-003.md#canonical-1010330312332021-0321321010002030-3101132022331133-3013020312210101-1111131022012101-0300130223020300-1000230210012222-2030012032210200) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--gcp_vpc_site--reference--group-003.md#canonical-0212103022310321-2033313012032233-0130103200332232-2313102132302210-1201322020113133-1301231313213311-1020223223233100-1120111212212022) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](resources--gcp_vpc_site--reference--group-003.md#canonical-1101300121301322-0303011310001213-2230223233212023-1002331032231131-0010202002221020-0210320200002221-1032220331202322-0000212030310122) |
| `ingress_egress_gw.outside_subnet` | [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-2122111320301002-2330002003210212-2333010033000011-0012210202210131-1130120220120033-3230212010123323-3302131321021330-0201230121321121) |
| `ingress_egress_gw.outside_subnet.existing_subnet` | [ingress_egress_gw.outside_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-1121233133132133-0330333020112111-2301122112021122-0323211003231111-0331232010210213-2212123021332003-3313112300320233-3020021022300130) |
| `ingress_egress_gw.outside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-0202000031230122-1000131131211032-2233100123323323-3221130023103110-0113312001300303-0323021222021103-2030213011033101-3213202122320010) |
| `ingress_egress_gw.outside_subnet.new_subnet` | [ingress_egress_gw.outside_subnet.new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-1231222031210310-0333220101322320-3310303121103111-2032130200211331-0333313333320302-0231110311023113-0122202311002221-1331022311330323) |
| `ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-1100310000332301-3310021011103222-3031032123233002-3213003212010320-3302221132232123-2112322331330312-2333023302011231-2222123311021101) |
| `ingress_egress_gw.outside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-3130022311233332-3230201100111021-0110031130122022-0110331023130321-2201031323033221-0130112320111320-0001300001121110-1000031021322110) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-1021120301130001-3022330333022200-3020322021220321-0321101003130102-3203100021330103-2213032202200130-0021000000002000-1320032122221022) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-1203013110311033-3120223331112012-2301113300131031-1022333313232200-0310023320031313-3001100230032120-0320023333302331-3023202310102121) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-2330013130311032-0332113321131202-3020112013103232-3013130331223133-3101230310030133-2232013033310221-0110133103103300-3123300013001231) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-2301133111200320-2213101300113333-1121211010300212-0121332231322020-0210322133321210-2001013312320120-3113033130231022-0313321300200010) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-3000132210103031-0220311220210231-3230012113110120-1020320123331231-0100121033200212-1130100223201100-1213211201102212-0201331300331223) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-1300011200010103-0333031112211210-0121212132313030-0023022313012333-2200113032121233-1011321100110320-3323110103103230-3331222230112010) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-0221301223203201-0132321123133313-0210320302230010-3211013020223110-2122311222012331-3322312320213123-0103100110211032-0111133233200131) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-3230222132301331-0301112123123333-2331212102323131-3220201301221302-3101212011102300-0233233322133103-1200030201120322-1201123302100301) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-1210230301210210-0103122211300023-0302003133222331-1311200021201030-1302123300331002-1130103323002201-0111313032210002-3300002220202133) |
| `ingress_gw` | [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-2121210113102303-2011303133203313-3210030103333203-3020313020302201-0111103112101112-3110132101132312-3133132320300133-3310212103131122) |
| `ingress_gw.gcp_certified_hw` | [ingress_gw.gcp_certified_hw](resources--gcp_vpc_site--reference--group-003.md#canonical-3222120301333200-0122001333021102-1310202303201212-3232212323321000-1113030032323001-1323000011230022-0210002322010000-0322303133210202) |
| `ingress_gw.gcp_zone_names` | [ingress_gw.gcp_zone_names](resources--gcp_vpc_site--reference--group-003.md#canonical-1231100301333013-2311222022012122-3131003132130021-2232020220003233-1320321320131133-2023322310002133-0210230311121121-1103133220232132) |
| `ingress_gw.local_network` | [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1112112002103311-1300002022011121-1130233321131032-0221232321213322-2322103021100300-2033130012022301-0333221033211120-0002121330221203) |
| `ingress_gw.local_network.existing_network` | [ingress_gw.local_network.existing_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1112220201022333-2203321013032030-0113221303033300-2023022123021311-3031020112200221-3230113303221113-2023321330320310-0030332013213223) |
| `ingress_gw.local_network.existing_network.name` | [ingress_gw.local_network.existing_network.name](resources--gcp_vpc_site--reference--group-003.md#canonical-3232020212100330-0331321120313200-3103232131103001-0020201312030221-3021333201120132-0333223120113131-0020102331023333-1130310322333010) |
| `ingress_gw.local_network.new_network` | [ingress_gw.local_network.new_network](resources--gcp_vpc_site--reference--group-003.md#canonical-2022301012320211-2123011330300230-1033000223302020-0022212311132101-1130111303123220-2320103313133021-0203331332112310-3320213231100210) |
| `ingress_gw.local_network.new_network.name` | [ingress_gw.local_network.new_network.name](resources--gcp_vpc_site--reference--group-003.md#canonical-1303313102032012-2101211221322301-1113031232220120-3201202230310332-2130131320310213-1203230320122302-1033203000013030-0232112231212313) |
| `ingress_gw.local_network.new_network_autogenerate` | [ingress_gw.local_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-003.md#canonical-3130210121003020-1312131003032000-0001113003203021-3213023000123113-3020212002012103-3031013113300301-2230330221223133-0312202202231133) |
| `ingress_gw.local_subnet` | [ingress_gw.local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-3311203330001001-3101330232102113-1233133230000131-1111003230221333-2321111323312302-1130023221302130-3302223022002013-0100123230122331) |
| `ingress_gw.local_subnet.existing_subnet` | [ingress_gw.local_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-1032003030331200-3201200021000320-3310220032022001-0132202031303112-3230230322211031-0132203121323111-0113311203030102-0033232310132133) |
| `ingress_gw.local_subnet.existing_subnet.subnet_name` | [ingress_gw.local_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-0102103202132310-1002300320310003-3002020030103310-3010232013033022-0200100121300301-3120100130102031-3232202213313302-0130110012022300) |
| `ingress_gw.local_subnet.new_subnet` | [ingress_gw.local_subnet.new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-2003133211221013-2221221012131330-0133000303101220-0322303003203311-3312102131113330-3232300203011310-0331302020013032-1212010131021323) |
| `ingress_gw.local_subnet.new_subnet.primary_ipv4` | [ingress_gw.local_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-3022120132320100-3322320102032322-1301302103232331-0002310333232123-2002233222101112-3211302203033220-0030030022331201-0010311203233002) |
| `ingress_gw.local_subnet.new_subnet.subnet_name` | [ingress_gw.local_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-0221032111100122-3321200212001131-2002001333200233-0121221020321322-1000221201230100-1130222131300133-2123113220111121-1213020021333120) |
| `ingress_gw.node_number` | [ingress_gw.node_number](resources--gcp_vpc_site--reference--group-003.md#canonical-2223330203110331-3220003021120123-3303201220320121-2310201033223202-2230313011332121-3302023011112013-1223220232100211-1210321213130233) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-2131330011023332-1203132212001002-3030321300303032-2211303212130122-2313001230220301-2031031213032120-3222001223233022-2110130302032331) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-0310313113011321-1022311211121132-0132000233001111-3123031331123101-3321002112103033-0233113311033120-1121302231233030-1301122101222222) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-3200120221201121-1001300012312022-0020212331100000-1121303131000331-2001210212100033-1020102320010323-3300323101202331-3211032233132132) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-1202222220310033-1000321032201220-1120003103022310-2203100322211332-3023102030231130-2222233123333323-2222331103033301-0010210312311002) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-2103130202233332-2202333020030210-3213121021031131-1223122130233123-2120211200110203-1223111131102121-1122233310313201-3331202000323032) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3232300231310013-3022111112222012-0212322010313231-0311223230313313-3103101121212301-3132230211111033-1223230213122202-2211121311011021) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-1202220301312112-0120003012301301-0133220132201123-3111323213311030-2100102313303330-0222331332203213-3121123003231211-0302021300332020) |
| `instance_type` | [instance_type](resources--gcp_vpc_site--reference--group-001.md#canonical-1122210310200102-2321323213130113-1032300030312331-1012022323131121-1012102233322210-3002311313033203-0231013301102331-2301023020133020) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-2201301300122332-1100020231023203-0132010132101112-1331031232112321-1123322231222102-0120332302321200-3201003130000330-1110011100002003) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-0320100202320301-3301201010000223-3210112200111331-2023121202111010-2332223121002021-3303133222132103-2101330313013131-0302300231303122) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-1110013123312303-0202130002130112-2110300232122212-3101132201211120-1213230102100101-1333210332102123-0133213321020331-3200303232303223) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-2303313222312000-1010211132211103-0102123312100011-0303023220311132-0221310233013303-0033121201310220-3310132201300323-2212332002313010) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--gcp_vpc_site--reference--group-003.md#canonical-0221122213321030-2312132011321120-3031022120023331-3232212011131022-0202123022132200-1203211132233200-1230100313033110-3210102313323133) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--gcp_vpc_site--reference--group-003.md#canonical-1301230203233121-2021332021030102-0213220001010022-0230130002220311-2033223022002022-1230112220303201-2032222021233212-1322331200210200) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--gcp_vpc_site--reference--group-003.md#canonical-0222122030033230-1032023002311320-1320300113200212-3303321333312312-2203300001123201-1310301201323222-2030202330100122-0322330022011312) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-3013200223033303-3020312130111111-2333110231223213-0202010110100333-2011232203212022-1133003132122002-3133020013120002-1203202210310213) |
| `labels` | [labels](resources--gcp_vpc_site--reference--group-001.md#canonical-2101102011223332-0323021220110131-0222123103022111-1003113313133332-0310332303232331-2210000012120013-1021330100101300-2103021222001310) |
| `log_receiver` | [log_receiver](resources--gcp_vpc_site--reference--group-003.md#canonical-2202123120111300-0113121321303331-2200201132003300-2210212100033330-3000030231231021-0201300131320000-3200201313121233-3303200333311021) |
| `log_receiver.name` | [log_receiver.name](resources--gcp_vpc_site--reference--group-003.md#canonical-3302331131022200-1300303322110111-0200033132203023-2032322000303303-2223302201131001-1212231122023220-3233210113033102-2313330133321133) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--gcp_vpc_site--reference--group-003.md#canonical-1332100133122320-1323102313201332-1331301323312130-2210131202023000-1130102331102023-3303311200001320-2333110320132232-2231033011302030) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--gcp_vpc_site--reference--group-003.md#canonical-3010202012213210-3023101201030211-2300101033021313-1301120111213222-0111002332203112-2210013220022232-3222311233103311-1232102001212013) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3320231131002111-2210100320021203-0222011020012223-1323133310102033-2211322122101220-0003233331201333-3201322031230300-1033023010213001) |
| `name` | [name](resources--gcp_vpc_site--reference--group-001.md#canonical-0322223203122102-1330212020100230-3212221100012111-2301110122322232-0130332002032012-1331301323223011-0230232130003222-3323002213300333) |
| `namespace` | [namespace](resources--gcp_vpc_site--reference--group-001.md#canonical-2002220012311032-2210330133321031-3310211030011211-3031013023203333-2302121121230031-2221032132101030-3313011130212201-1322101232301122) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-0202021003300112-2312231030303221-1030200233033323-3213311032213012-2013100031301112-2023120031120303-3230133013302032-1323323332103110) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-0202212320213211-2031130231021003-0323230030112230-2013002131123003-0131011211212110-3010212212020333-1333022312233300-3332321301200311) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-1021213100110120-1011030120111131-2013003010103210-3113100311302102-2231010100133313-3112133112013223-2002011013130212-0103102301012120) |
| `os` | [os](resources--gcp_vpc_site--reference--group-003.md#canonical-2020323220131321-1201222201012212-0110113022333313-1011013303030202-0213332100122103-3030323112022220-0123303300231302-0203211313310020) |
| `os.default_os_version` | [os.default_os_version](resources--gcp_vpc_site--reference--group-003.md#canonical-3031223321221020-2031131013001211-3000320223320221-2333203103323311-0322103301011123-3002230121331331-1133131022201010-0132122012110001) |
| `os.operating_system_version` | [os.operating_system_version](resources--gcp_vpc_site--reference--group-003.md#canonical-2111223213022002-0313302330013333-0032231322020212-1213330011123133-3031312022310110-2003302122200110-2010210131232021-2223001210103001) |
| `private_connect_disabled` | [private_connect_disabled](resources--gcp_vpc_site--reference--group-004.md#canonical-2233301233101133-1320123210011220-2031030310331002-1213231021002033-3120333011113032-0311131033130112-3131132121200013-1113011101123023) |
| `private_connectivity` | [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0012302121203202-1121213313020210-2231002021013133-3303022320311132-2003013103011130-2102110312212203-2330210031022200-1121300300012103) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--gcp_vpc_site--reference--group-004.md#canonical-0210000023220202-0033032230123002-2020021031131221-2032101233131303-2013103113010221-3231100313213003-2210113002220211-1022022303333032) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--gcp_vpc_site--reference--group-004.md#canonical-3323232130122201-0311200223131222-3110030001033331-3032023131020330-0201003123322113-1323230020230332-1102232111232212-2302323011232202) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-1022031331201123-3200311022312321-0120200232332201-0231201331331003-2233023120202310-1012320012300200-3322111311113013-0321310233023230) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-0210321321210230-0200000322131012-0221132101112320-1110033313200231-0111322000330021-1123002212223231-1033020101120100-1321030011230133) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--gcp_vpc_site--reference--group-004.md#canonical-3023020310313033-1133101321220232-0002332000020001-2322313121100320-3130212313130331-3021003130003310-1111000022132222-2123310211310022) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--gcp_vpc_site--reference--group-004.md#canonical-2301130220031333-1032302232212020-2000130322333312-2000000210223022-1132220030211001-1030320301212230-2101123221122010-0212102311312302) |
| `ssh_key` | [ssh_key](resources--gcp_vpc_site--reference--group-001.md#canonical-2221333211121223-3100331330030210-1120033101130210-2003113111111330-3023332233122030-3200210331311003-2121310233223101-0032123133313012) |
| `sw` | [sw](resources--gcp_vpc_site--reference--group-004.md#canonical-2230022121132012-3223032230323023-1320010311112313-3023032010031002-1023002302330022-3231111312200112-2200100031111003-3210133123231221) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--gcp_vpc_site--reference--group-004.md#canonical-1022302300321021-1312023001231313-3313201023130130-0201122333001310-1332113111233313-0130213323102132-2000023301102210-2313212211223202) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--gcp_vpc_site--reference--group-004.md#canonical-3123111223133002-3012011121232301-1333301202213200-0302031301221220-1200030110222023-3202130212023012-1133221022000132-3231002223021122) |
| `timeouts` | [timeouts](resources--gcp_vpc_site--reference--group-004.md#canonical-2032223202203233-0230223113122212-0211131303203223-1211220330323100-0233000102221303-1010131310223011-3112011211100110-3210300013032132) |
| `timeouts.create` | [timeouts.create](resources--gcp_vpc_site--reference--group-004.md#canonical-0320201010322333-2231210101121230-0022130122331330-2230131332003111-0120212302030021-0331333022020321-1312201000131233-3323212131323033) |
| `timeouts.delete` | [timeouts.delete](resources--gcp_vpc_site--reference--group-004.md#canonical-3020113212011120-3001313311333111-0321223123123312-0233013212302221-0213211002001013-2012100230021120-0023122022233310-2332231203200023) |
| `timeouts.read` | [timeouts.read](resources--gcp_vpc_site--reference--group-004.md#canonical-0312331023101330-2133031121203110-3330010302133111-1321112020013322-1132323002201213-2112123211213311-3000020300102123-0201332221303310) |
| `timeouts.update` | [timeouts.update](resources--gcp_vpc_site--reference--group-004.md#canonical-0133213020320121-2110123122032222-1230311110330232-2332022231233101-2023332111322113-2210332012130112-0002113311103210-2332113302001212) |
| `voltstack_cluster` | [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-1312003312232212-2011113310110131-0010131322113332-3131330000201333-3320312120210223-3032202123102121-1200233120130210-3103303212322322) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0323220302231320-0232121331330322-3320303232313213-0202233102310123-1323332032313110-2220323202020133-3113012302012012-1212020202333333) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0122121330302212-2112001103023230-1131102223123110-2023333210321233-2010230133102320-2313030021213233-1011213012210322-0200322333032313) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--gcp_vpc_site--reference--group-004.md#canonical-3032232303130002-3233231133013133-1212010220221220-3031121202032100-3003201332233103-1333133131311223-1230321012010221-1013201123323100) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-1121020223121322-1331021002022321-0023013330031133-0121013330121313-1313210322220123-3223103121031033-0330313002123120-2130123011032003) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-0222211103322213-0331130333323301-2021220300113101-0331121011310030-0112212010332120-3023230101000031-0021320103020302-2310312322010123) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0311011300332003-3302003020211322-2321323021121001-2321311220112122-0223333001031220-0133332332110312-2022020323330003-2212001322312331) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-1213012333101200-2003333223031010-2001323131021000-0312011102013000-2220330020033121-2303010202232230-1223020103301300-2220131332033201) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](resources--gcp_vpc_site--reference--group-004.md#canonical-0113103220223013-0111231033333130-1330222312020302-2131331002021223-2222320232302313-0330011320011101-2120122032313233-2130311021300231) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-3021003223131231-2220032310033113-3133202111221220-1120203111230112-3203030003101030-3122011133302111-2301020303232322-1331121312201333) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-1332021102110202-1230001033223202-0200020210201220-3230003303123131-1121302030221021-1021102322023202-2311332331220121-0330323000110011) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-2022130321323021-2102210321321230-2030310212001022-3300332330032231-0201112031333301-3102000111132020-2233132103303032-2010012033313021) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-2121121000131320-1003212221220303-0333332120202223-1101113230031311-3311131231211231-0003300202001121-1222021232330101-3201202002300112) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](resources--gcp_vpc_site--reference--group-004.md#canonical-1201230223321221-0033131311132102-3310111301301331-1122031121011330-2332313101113023-3011323232201100-3022303203013310-2312211001210231) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-0011222030213233-0030023021120331-3113212031321221-0130131022133111-0103221121331200-2322131000011311-3322202330133033-1301232202223302) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-1313200332231031-2233021131211132-0020212223123202-0320011012100100-3210230301103001-2232302030311033-2230001321220221-2200122000020231) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-1011210220301302-3213010231210001-3011201022223230-2103011330131201-3302330222023010-0021111231102223-2303001132131020-3102011203221030) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](resources--gcp_vpc_site--reference--group-004.md#canonical-1333210110211132-1330233030111012-1023100320203231-3000031201201302-2103032201332311-3103111031011023-0131123131000101-1023322211013330) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-1101132001130021-1313322101030332-1032313023010313-3302330131111200-3011122220220321-1123133122333000-3233220003023303-3221023020011000) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-1303021133231322-2201113001330223-3223203330223330-0001031020303030-2130122330031332-1000200232020212-2022103030310210-0232311212012203) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](resources--gcp_vpc_site--reference--group-004.md#canonical-3210210310130012-3131230000101321-3223100200303111-3211000132123211-0013133300323303-1230023330030001-3011220232131102-3303310001310113) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-004.md#canonical-2302130023333312-1131022320130133-3111001022331233-1222310232112122-2033120323220102-2111212331011023-0000220012100323-2222113320323302) |
| `voltstack_cluster.gcp_certified_hw` | [voltstack_cluster.gcp_certified_hw](resources--gcp_vpc_site--reference--group-004.md#canonical-1230032001131031-1211111322120131-1103200133101021-3000233300301303-3122330030102302-2100202320120230-3103230020133000-1220300331220200) |
| `voltstack_cluster.gcp_zone_names` | [voltstack_cluster.gcp_zone_names](resources--gcp_vpc_site--reference--group-004.md#canonical-2032023210311213-2321323333200032-3223210002113301-2310212320232133-0023000000313203-3111200321332111-0032003220231122-1122302020321001) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-0211002121113202-2211002211223231-3030300202113003-0221001022123333-1103210312332321-2323321331200131-0221331310031331-0023101131120110) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-1131312233313233-3131120211100310-0300223313213031-0331331323022120-1021321330133130-3103233222003111-1212030222121003-3222022020031103) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-0221012313020012-0230333232210333-3211201322332301-0220201212323313-3322201330001020-1331321203313000-1000001213001323-1022031001010223) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-3010002033200322-2201200101303221-0213230322300033-2222033323001211-0021021010121311-0110113022211221-3031133131033121-2220110323212312) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-004.md#canonical-0002311201021301-0122021323123213-1232211012331330-1121332323001020-3302200313002002-1113112320233130-1332333230222212-0112112212211223) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-3001101132320132-2220330330033020-0330130030001023-2021032331132111-0033003331230130-0003310321310221-2223233221102133-1131322121212203) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-1112030222023331-3211101220032100-2031330202210020-0323231333212203-3230120213300230-2112123033210031-3322012201021113-0122210031201121) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-3100001031330120-3201033101033122-2001011132322031-2021031200200020-0033002121222221-3102333203130232-3133231200301030-3202332310331210) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-1303102012101312-3313102330012020-3002310303123303-1211010220222212-1232000323013322-3212302110001333-0210031333311302-0000011022032031) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-004.md#canonical-1231023201212121-1030120002120322-0031202131102203-3032303012111111-1211312013212113-0312103231323101-0220002123110313-1310232231212130) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-1313223202322311-1020202131332120-1313133130322002-2130330212302201-2131203022100002-3000101122100120-2320112021231322-0101221030322202) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-1223230113233022-2231322031000002-0003102013311011-0131031111030000-1303112220300302-2130303311322002-3032003120100210-0022313213030130) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0223331013200000-3123101202112122-3111022230211032-2103012130110220-0200332103130310-1030001330131212-0310023032123110-1113330203222112) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](resources--gcp_vpc_site--reference--group-004.md#canonical-3133302130131000-0121230133122000-0221201103130132-3222010231130000-0001321323112121-3132113031111201-2313102232021031-1123213301123123) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-0221210111323232-0221123303101100-1000311233312002-3113121023320131-1032030121001033-3303323301113202-3121021010023310-2032322110010120) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-1102312111120301-0312332031320002-1222201331132310-3323202003032001-1200112331331031-0111030000330233-2321101332012130-2031233203020123) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-3220201320120000-1011012212201120-0310330213001223-0210320030203313-3211231132013303-2211210002203233-3010322020031222-0321110112101112) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](resources--gcp_vpc_site--reference--group-004.md#canonical-2120011113003203-1101121111012011-0113133322300021-1022331003233000-1202110301031300-1211321302032012-1330310111133101-3311100211330022) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031301202330333-3002113133000202-0313100033022232-3133200011301112-1302220032203030-1211231211212300-3232022003123020-1332122211012313) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-1122212012233000-2223100212222031-3021033010301210-0333123130211002-0312301132120222-3201101303002110-2032331121221310-2100122031131102) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](resources--gcp_vpc_site--reference--group-004.md#canonical-2320212130100013-1030210112323321-0112102033333200-0301003023021003-0230122210312201-2221000100211301-2300313322030233-2101223123021130) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-0103131130011222-3100230002202330-1311223202303203-2112031022122002-3033013313030021-2200201201232302-0202120113201123-1232033333313012) |
| `voltstack_cluster.node_number` | [voltstack_cluster.node_number](resources--gcp_vpc_site--reference--group-004.md#canonical-0030103222301332-3301301102132020-3303110220010031-0313133030233010-1201321330021002-2300202110332000-3110012012203000-2000100232030101) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-3111233033111112-2103121333111022-2111001110303210-1132112013331012-1011030332211001-0200213031001200-3301111002301222-0021120011123301) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-3002133220021122-2232023003113121-3322333222302220-2231233310330022-1103321233001111-0102102300200030-2211233303330010-2130310231303023) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-1121310230231302-1303033113022033-2113202202320303-1103011331332213-3103123130030100-0020223121320322-3133020113212320-3021331323032311) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](resources--gcp_vpc_site--reference--group-004.md#canonical-1321002003333012-0010002320333322-2011312332332311-3122132202221001-3212120201213010-1321131233313231-1030221121220233-1322011110011321) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-004.md#canonical-3101111101022131-1111200112303320-1311303001200023-3021200301211001-2020133320020120-1300030122100303-3232332100300322-3022213113231110) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-2232231233200232-3100310330201323-1003131223003320-1313222222020333-2203010203020011-3223022120313121-3011100020220221-2013002020223133) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-004.md#canonical-0302101032201101-1302022233021030-1012012122031020-2112312201132110-3231103013323001-3130111122122121-1132001101131313-3100233220010111) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--gcp_vpc_site--reference--group-004.md#canonical-3222101001203201-3131211303032022-2301123122313031-1201233100121322-2102223333012131-0021210303323333-0230200200301311-0002220103001010) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--gcp_vpc_site--reference--group-004.md#canonical-3010212233232231-3013033113003123-1203302032231301-1211300120110333-2323300102031032-3231231033020320-0013222100030132-1233012301122312) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-0110310300111123-0132301311011212-2110322031132312-0220212322010131-1330013112301023-3211020310123312-1121301120022110-3333210131311130) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-0021211213223331-1131313201023311-1233133201323313-3130303321320210-0031120302212000-3300233030332022-1311202312232010-3223331312121302) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--gcp_vpc_site--reference--group-004.md#canonical-3020223103022331-2312033002231032-3122311232212133-3001330033023003-0013320102111010-3220331331303223-3131321211012220-2311112202221033) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-1121311211102200-2020311221021133-2113313203321020-3322231333113123-2022212301303310-0233310111123002-1302101001212323-1030201022123303) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-0222021023021233-1023112033000200-0002101112022331-0013322002033010-2102231032002302-2113331322211211-0022332010023133-0323201123331332) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-3033120030123232-0002233000313233-1111222110330301-2121231333203002-0202132132233110-3110132133020222-0001213331012010-1322133321220110) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-2200222213312102-1032003201302310-0103302233312300-0133120001122222-0223133130333120-1201302101303132-2232023330230000-3022100202032012) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0221133130132033-3012120122302320-2121300212022202-3211130110311120-3332121300023333-1010201230112121-1202230220110122-0231101233113130) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-1020103130332111-0303220101020320-1333221100102232-1310212032022311-2302112222231220-2213131111012023-3310322010211211-3121122323021132) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-0101133221333000-2220313031102310-0300030222131212-1310321101330102-3101101001112212-0312111132020210-3301300133320303-2012301111212323) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-2230230022133221-3113330230301232-1020320322230032-3003030122020212-0323201330231320-0233113033113133-3300222312132112-3001302320222310) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0122013301300302-0310321332303333-2323002030332131-3021112110130201-2131302213131210-3120202110302010-3322102112032003-3121121233120030) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-1033321000122221-3113021022332013-3212022100232111-1032232302331123-3031232131010322-3330330123223012-3020020211123330-3121301313132023) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--gcp_vpc_site--reference--group-004.md#canonical-0303112110013220-0311032211310332-2011131111010111-3110121130310010-3230333122231121-2310103030113222-3011221311033302-0030210222302110) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-0313100012010200-0120320212213231-0233110212321331-3010312300011033-2113122200310223-3013030103203320-3122230102033122-3232030132220201) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-1231210021323323-2310120233230322-1130030110123333-0012111313333030-3121212201003201-2110122022102311-3013122331302320-0320120322320102) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--gcp_vpc_site--reference--group-004.md#canonical-2132101120321210-3232111120230021-2101013112020220-0020112022000102-0022223210022022-3232031122200202-1113301302231110-0320211203302312) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--gcp_vpc_site--reference--group-004.md#canonical-2121102320210200-2023123101312300-1123330022021312-0222223130103311-3131121020112303-3200001100132323-1320321212211012-2220322323133203) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-1123013323112120-1101200003032122-3121200200213203-1221232303201320-2200111323101223-3033121220313102-2302021123130330-2030203112233220) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--gcp_vpc_site--reference--group-004.md#canonical-3120211330030011-2102113122220312-0100332133230112-0222133321103130-1001023101200113-2222102023111210-2303213032223312-0333003303021020) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--gcp_vpc_site--reference--group-004.md#canonical-3111223132123130-2213001312220122-3331031333133221-3312330032132310-2132021231032132-2333000112021200-3002132122102302-0313222010310203) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0330102112000330-1030023011222100-0310320100211300-1201320202321323-1032112100220011-2021310102212121-1203030132212333-1120201210232301) |
| `voltstack_cluster.site_local_network` | [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-3312112033010321-3313133200100220-1312113103010000-3012203331320010-3033332300011210-1303010331223000-3310003200132201-0210013203202012) |
| `voltstack_cluster.site_local_network.existing_network` | [voltstack_cluster.site_local_network.existing_network](resources--gcp_vpc_site--reference--group-004.md#canonical-1303001131232032-3302112102322133-3311132233000001-2031232010012022-1213202312232100-2030312101210130-0200311212032100-0331102203020002) |
| `voltstack_cluster.site_local_network.existing_network.name` | [voltstack_cluster.site_local_network.existing_network.name](resources--gcp_vpc_site--reference--group-004.md#canonical-3010302131010323-0123310022311122-3311203101221123-3312320022202003-1131313203233000-0132013233121323-0002031021200122-3113020311130112) |
| `voltstack_cluster.site_local_network.new_network` | [voltstack_cluster.site_local_network.new_network](resources--gcp_vpc_site--reference--group-005.md#canonical-0203112011300100-2232221113202223-1002330210332203-2323131202103001-2221321201022013-1200123201221112-3201302030013331-0310330102320111) |
| `voltstack_cluster.site_local_network.new_network.name` | [voltstack_cluster.site_local_network.new_network.name](resources--gcp_vpc_site--reference--group-005.md#canonical-0331302222121311-0303133100223212-0330030202332101-3202013320130222-2330020123013102-0032323300321320-1212122330231330-3030131231211123) |
| `voltstack_cluster.site_local_network.new_network_autogenerate` | [voltstack_cluster.site_local_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-005.md#canonical-0003300001200311-3103132200321101-3312300211211302-0010203021333111-1010132023120233-2001001303331310-3301300111322331-0331020313232300) |
| `voltstack_cluster.site_local_subnet` | [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1112233310222231-3323120000012103-3331013012230112-2010330332111120-1332030031220222-3312133112303103-1200201133331310-2102020021121331) |
| `voltstack_cluster.site_local_subnet.existing_subnet` | [voltstack_cluster.site_local_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-0010213320213230-1200122022123022-0013100100123331-1102011121101112-3301130302333201-0023203120131302-2103221301101013-1011330133131230) |
| `voltstack_cluster.site_local_subnet.existing_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-005.md#canonical-2110213000321012-3322030300311112-0223230113003003-3123232330312010-0232312320221233-1032321311333231-1211032303003121-1003033121300123) |
| `voltstack_cluster.site_local_subnet.new_subnet` | [voltstack_cluster.site_local_subnet.new_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-2333203110321011-3221113220130113-3223003100020201-2030130112012202-3212130130113220-1201111131012022-2112302200310330-2230313121002000) |
| `voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4` | [voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-005.md#canonical-0121313112032323-1030012112032203-3211102210303310-2033332323103111-3020203113012112-0122230200212300-1312001200303120-3212002132311002) |
| `voltstack_cluster.site_local_subnet.new_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-005.md#canonical-1022333112000021-3223132022301030-2213100133322122-2032121110002003-0323221310013213-0110101001132002-1322232120030212-2301022300123012) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-005.md#canonical-1330023103101131-1323220103232030-3323301010002103-2210330220321133-2110031213333013-1010031200101022-3202033111010302-3203311032110223) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-005.md#canonical-0001231201231133-1001001202203233-3311113131230131-1221031000210331-2301230312233001-3102033110232132-0331011113322022-1211011311231002) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-005.md#canonical-2012112022110310-3212303113133310-2000330201133312-1111100233121020-2021201323101033-1001130210331301-2123122110012312-0013212312300321) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](resources--gcp_vpc_site--reference--group-005.md#canonical-1233311323031330-3111022113232202-0221021202303321-2210300122211231-0030023331022333-3131120113202002-3122001201330231-1120120010112003) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](resources--gcp_vpc_site--reference--group-005.md#canonical-1013230323222321-1201113220332032-3000020121233231-1333211000101323-1100333310201212-0321311231222121-3123100011120320-2231331212213003) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](resources--gcp_vpc_site--reference--group-005.md#canonical-2031121313222202-3123233001232312-3223303212130300-0212223230011330-1003132330201221-3202113320222202-3022320302123020-3213221113033210) |
| `waf_signatures` | [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-0200303132233121-2101103333000100-0121030102330013-2121200130113333-1302301003223332-2320122122320002-2210031023012000-1021113133221212) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--gcp_vpc_site--reference--group-005.md#canonical-1321103100021330-1221211333223021-1203223102233311-0102130330003133-2130331112110030-3131033032210320-2033332233131233-0003103320233211) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--gcp_vpc_site--reference--group-005.md#canonical-3301321103032221-0303320300031322-2111301110210302-1331000211302200-3323323133321311-3002113320123013-2322303002133100-1230010010130103) |

<a id="canonical-2322110222223231-1113312032301332-0010212112310220-1210231031031021-2023203332032103-3311200313232220-2012023112011323-3003133220302121"></a>

## Next pages — Property reference / 020322303323 / 18

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-0120223131022023-3232101001310331-1103010012022230-1200131021301230-1312302111223330-3010233323332131-1111303112300132-1103200021202303)
- [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-1103003322103331-3002300011231332-3102223020003212-3211211233001332-2032120010123310-0313231000122102-0303201323233323-3101332311030331)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032)
- [cloud_credentials](resources--gcp_vpc_site--reference--group-001.md#canonical-1312001102001300-1311230131021121-2213102022222331-2331123300121313-1323121213000112-0102121102020303-2231311102230221-1322021220133331)
- [coordinates](resources--gcp_vpc_site--reference--group-001.md#canonical-0132322102032221-0221031120301201-1222222320102200-3032133011203110-3012200020001320-2313113033220032-2222201130301012-2121000101223020)
- [custom_dns](resources--gcp_vpc_site--reference--group-001.md#canonical-3220300023200020-3013222211030303-0100203211331300-0030113331233123-3331132002222231-0000301012011231-2310300112031020-2032323332321030)
- [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-3310031330331201-2123120101221211-0113032011130223-0003312130311301-1303121010101221-3233302131102330-2311002110111331-1013112312302133)
- [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1131111322332132-3103020101000022-2211033321232233-2030233213233230-0131222130003033-0220320200112322-2313123101130311-0133321200322303)
- [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1101221123132310-1010130003200313-0032320202321131-0203333233313013-2011212112020022-2213003211113113-1032113120210301-2132310133203302)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-2221222220302122-3232110120322221-3100201231013000-1213023302033003-1313222101212322-2133001213332231-0312122202112110-2121222013023030)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-2101212102330031-0010112330122012-1102313232223312-0222112303201021-3012113311101013-3130313120212310-1030123002232113-0300133312323323)
- [log_receiver](resources--gcp_vpc_site--reference--group-003.md#canonical-1023013120132311-1101001323331302-1302211232003313-1023011331032223-1301113322212223-1030321303032302-2211010200203113-1223301000320112)
- [logs_streaming_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3001101023203132-2313313210112200-0002300333010120-1210100002212200-1202021011233220-1222101232112112-2003010130010021-0013120113331012)
- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-0302011120033022-2021220221211032-3112222301002021-1021012200222002-0322233013032322-0132210231022013-1300001101112021-3332323033020031)
- [os](resources--gcp_vpc_site--reference--group-003.md#canonical-3222223210301122-0312211222330230-3122221002123233-0113101130301311-3320303012132320-0302323231130332-0123201033332330-2211021130022023)
- [private_connect_disabled](resources--gcp_vpc_site--reference--group-004.md#canonical-1200311111133100-0020033221223212-3021002202113332-2301121021212132-0232212133303310-2231010113102221-0033221003131012-2022210321113201)
- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133)
- [sw](resources--gcp_vpc_site--reference--group-004.md#canonical-0212133321121330-1301131103023331-2123132213311103-0031232100002231-1233033231110220-1102201033230030-3020012013122301-2110012031323002)
- [timeouts](resources--gcp_vpc_site--reference--group-004.md#canonical-0202212023111322-0302222202030003-3121101231202200-2213031121020101-0121233300121223-1032213023303312-0220022023213022-0231230022223121)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-2220233032031010-1323100300312032-0222201202010130-3130203022232222-0300133130321233-1132212212212321-0322012021310311-1132132110223100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0120223131022023-3232101001310331-1103010012022230-1200131021301230-1312302111223330-3010233323332131-1111303112300132-1103200021202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232100030131212-3231323103112112-1202332312323012-2231132320023130-2212202110311223-1020133232313232-2013303112221002-2021230202000310"></a>

## admin_password — admin_password / 021221301031 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- admin_password

<a id="canonical-3003220023201110-0123012222203123-1123232121100300-3130110132221221-0003313012032303-3232221111101211-1120011000101222-1122023112003130"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322032333022321-1020321321113003-1032301132220102-2221320321323130-3110111301122301-2133130311301033-2333321320333223-1332002113321200"></a>

## Direct properties — admin_password / 021221301031 / 3

- [blindfold_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-2111222221030100-2021012131222123-0321031131332322-2012321203023023-2331100121000210-2123120220110100-0313002312011332-3320313131032100): complete subsection reference.

- [clear_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-1203200211201210-3332131011320312-1221211301322322-1112202232332120-2310013330323122-3323320022110330-0101232201221010-0220300120010100): complete subsection reference.

<a id="canonical-2310320230013233-1120221013101311-3003311113301301-3132120021122222-1121312111322000-2300330122132101-3332321202210230-3132213202322101"></a>

## Next pages — admin_password / 021221301031 / 4

- [admin_password.blindfold_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-2111222221030100-2021012131222123-0321031131332322-2012321203023023-2331100121000210-2123120220110100-0313002312011332-3320313131032100)
- [admin_password.clear_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-1203200211201210-3332131011320312-1221211301322322-1112202232332120-2310013330323122-3323320022110330-0101232201221010-0220300120010100)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2111222221030100-2021012131222123-0321031131332322-2012321203023023-2331100121000210-2123120220110100-0313002312011332-3320313131032100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320013011210111-0332021002200221-1210231032330230-0012011223212130-3202122123002223-0221103103221200-0203113320033031-0010000300123020"></a>

## admin_password.blindfold_secret_info — blindfold_secret_info / 333120003210 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-0120223131022023-3232101001310331-1103010012022230-1200131021301230-1312302111223330-3010233323332131-1111303112300132-1103200021202303)
- admin_password.blindfold_secret_info

<a id="canonical-1332023311103202-2230103223031121-2121233332332103-1223132013113101-3311120330022011-1000121101313012-2211023022132033-1303032222130010"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131303303101220-2022330103333021-3113322000313321-2313112300100230-2333012030033211-3322100202212030-0012101330031130-3012212011011122"></a>

## Direct properties — blindfold_secret_info / 333120003210 / 3

<a id="canonical-1312023222220223-1012130203310030-2013313201332031-3003013233310030-2230110002211332-3201003311300111-2002001221021300-1111132213211100"></a>

<a id="canonical-3110021203133333-2000311022022023-0122011030032300-1301011332331331-2201301011022133-3103000031012101-1030321231310112-2122012223123103"></a>

## decryption_provider property — blindfold_secret_info / 333120003210 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0312331110211323-2320303321232213-0120201001023022-3002220133003030-2132221133130002-2101110121021311-1232210312301322-0203023001222333"></a>

<a id="canonical-1331123000202300-3103103231121201-3010322133000233-0300313232112130-3120022313231122-0123010302312103-2310032111100001-2221333103032113"></a>

## location property — blindfold_secret_info / 333120003210 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3233031121121033-3213233102322320-2203132020211022-2133020100311321-3331331213120233-1010111200211113-3230000330030120-0100123031032012"></a>

<a id="canonical-1133310002210030-3020032201303333-3312002203311310-3131310131201212-1021103213113130-3213322210021020-0203001020023022-3223120123031302"></a>

## store_provider property — blindfold_secret_info / 333120003210 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2122213223210011-3211112013201011-3302201112111230-0023201111312323-0113303112121102-0100211033311133-3110301022113302-0231223002130033"></a>

## Next pages — blindfold_secret_info / 333120003210 / 7

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-0120223131022023-3232101001310331-1103010012022230-1200131021301230-1312302111223330-3010233323332131-1111303112300132-1103200021202303)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1203200211201210-3332131011320312-1221211301322322-1112202232332120-2310013330323122-3323320022110330-0101232201221010-0220300120010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032200300132122-3322333021003200-1310333031113302-3012013121211310-0322330301102013-3202001231311020-0203320013203203-3310222333113110"></a>

## admin_password.clear_secret_info — clear_secret_info / 200232301312 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-0120223131022023-3232101001310331-1103010012022230-1200131021301230-1312302111223330-3010233323332131-1111303112300132-1103200021202303)
- admin_password.clear_secret_info

<a id="canonical-0023012333203103-3202010223113032-1312333031222113-0030203120110101-1222113222031302-2131202332303222-3311101331332030-1023012321301102"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330022313111312-3033101002200320-1300121331011320-1321321031310121-1012202320031102-1003003210321200-1122113121322111-0230322300123231"></a>

## Direct properties — clear_secret_info / 200232301312 / 3

<a id="canonical-3113320223122002-1132200102120012-1022212120000302-3030022130123131-2101231303213110-2103332021313330-3223311112212332-3201110330232133"></a>

<a id="canonical-1232123003033310-1122323123113322-0021121101200012-1112031100000121-2111221213300010-1102112022220202-0102131111120100-2303333313222012"></a>

## provider_ref property — clear_secret_info / 200232301312 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0323200001031310-2121103020112112-2201003200032121-0002002310032001-3221003112230320-3103312200121130-1302311022133022-0210123001020002"></a>

<a id="canonical-3101300222323123-1332211310201211-3012130321233112-1230122323301222-1030221223221203-0202220001110333-0331132302231021-3010102211123212"></a>

## URL property — clear_secret_info / 200232301312 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1323332110213202-0013320130111021-0011300011332032-1102123012100102-3031103222132022-2320212031313100-2001112331123032-3003221331123233"></a>

## Next pages — clear_secret_info / 200232301312 / 6

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-0120223131022023-3232101001310331-1103010012022230-1200131021301230-1312302111223330-3010233323332131-1111303112300132-1103200021202303)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1103003322103331-3002300011231332-3102223020003212-3211211233001332-2032120010123310-0313231000122102-0303201323233323-3101332311030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221200121313321-0112222201332113-1203321101120030-0230232130113311-2323311032302022-0133112211122323-0111233231302001-1112322000200313"></a>

## block_all_services — block_all_services / 332301213203 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- block_all_services

<a id="canonical-0010203330220110-1321221020003123-0313333201303011-3320312311130233-3333101130013200-0111022232312012-1302103101110030-2223213332030112"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-0010203330220110-1321221020003123-0313333201303011-3320312311130233-3333101130013200-0111022232312012-1302103101110030-2223213332030112)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-1000323100231311-2301320123232312-2002123332011201-0233121312022123-1011032320123203-1103333123203103-3003223032211302-3313232121221332)
- [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-3221320222002002-3132322013032031-1231302031321113-3320022220113332-3202301310311303-1213132233212112-2000011221113330-1331212333223123)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

<a id="canonical-3022123010231221-0311301021210102-3330002021220303-2103332013202020-1322013000222020-1233122021023000-3000103123113313-0013321023231213"></a>

## Direct properties — block_all_services / 332301213203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131300213013321-2320202222233230-2033332013012303-2121302232032323-0202330330223232-0210111032321103-2222303211223202-0132130101013103"></a>

## Next pages — block_all_services / 332301213203 / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131311233311332-3223133302203133-1013223121320232-2012233101233023-0221012301211331-0000213230002103-3221123202332200-3301312130030302"></a>

## blocked_services — blocked_services / 031230010202 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- blocked_services

<a id="canonical-1000323100231311-2301320123232312-2002123332011201-0233121312022123-1011032320123203-1103333123203103-3003223032211302-3313232121221332"></a>

Type: `"object"`. single nested block, Optional.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103003300113103-2101131232220213-2323120132323331-3210123311331202-2021002030202301-0113030121323021-3222221311020202-1333220323133132"></a>

## Direct properties — blocked_services / 031230010202 / 3

- [blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200): complete subsection reference.

<a id="canonical-2202131103022110-0232120132100130-0310132303331022-3013101211123201-3323113330313000-0200232303020321-3332030131302103-3311010212321022"></a>

## Next pages — blocked_services / 031230010202 / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222002311130333-3331332222003002-3311101100321330-2323001210202113-2110300332031113-3313312300230120-0102100002101213-0113021030001112"></a>

## blocked_services.blocked_service — blocked_service / 002323121012 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032)
- blocked_services.blocked_service

<a id="canonical-3032113103131331-2020313011111333-1023323013120331-2300012103020012-1113012330021330-0132321102332203-1300012122032222-3123330200221010"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301312320231100-3321313332321213-0032222221023313-0203323212131213-3130121122022022-3321230003313213-2210313232310231-2021202112320122"></a>

## Direct properties — blocked_service / 002323121012 / 3

- [DNS](resources--gcp_vpc_site--reference--group-001.md#canonical-1213200203033320-0113311321120101-0022311002102301-0332021013231003-0022300211220222-3230103232321132-1231220222022313-0232321130121013): complete subsection reference.

<a id="canonical-0331102313120332-2313220333130202-3230313311321311-0012332232321210-2321231130001212-2301320301310032-3222032122202002-2231200031101020"></a>

<a id="canonical-1112100133303223-3221111103130032-2010310021231032-1003310022110331-2310000120120100-1002303113201030-1202211202023110-1312032031301332"></a>

## network_type property — blocked_service / 002323121012 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [SSH](resources--gcp_vpc_site--reference--group-001.md#canonical-0023321322011323-3100133130131010-2130132310133012-0320103021203132-3302111230110322-3233011130100101-0303133131312311-3033112213223210): complete subsection reference.

- [web_user_interface](resources--gcp_vpc_site--reference--group-001.md#canonical-3010321200301213-1221323211101322-2333032333112021-3301003333002322-1213032203210021-3210301202210213-3122033232331203-2020312120222230): complete subsection reference.

<a id="canonical-2313112103010222-0312020111302332-1020232323323231-1200203033223312-3123030200013203-0112231333002012-1032221032011130-2120122233023022"></a>

## Next pages — blocked_service / 002323121012 / 5

- [blocked_services.blocked_service.dns](resources--gcp_vpc_site--reference--group-001.md#canonical-1213200203033320-0113311321120101-0022311002102301-0332021013231003-0022300211220222-3230103232321132-1231220222022313-0232321130121013)
- [blocked_services.blocked_service.ssh](resources--gcp_vpc_site--reference--group-001.md#canonical-0023321322011323-3100133130131010-2130132310133012-0320103021203132-3302111230110322-3233011130100101-0303133131312311-3033112213223210)
- [blocked_services.blocked_service.web_user_interface](resources--gcp_vpc_site--reference--group-001.md#canonical-3010321200301213-1221323211101322-2333032333112021-3301003333002322-1213032203210021-3210301202210213-3122033232331203-2020312120222230)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1213200203033320-0113311321120101-0022311002102301-0332021013231003-0022300211220222-3230103232321132-1231220222022313-0232321130121013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313332010231311-1320001033120311-3000110301313020-1333321221020223-1303003232033013-0010023301021302-2231223233133321-3023230131132010"></a>

## blocked_services.blocked_service.DNS — DNS / 011002303000 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032)
- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200)
- blocked_services.blocked_service.DNS

<a id="canonical-1112230303310220-0031101202233200-0201131111111021-2030022133021000-0223013233131031-3321202203222110-3002033332230231-1121031122113330"></a>

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
dns = {}
```

<a id="canonical-0333221302101021-3103022221113131-2000032233031110-3223231223300121-0200312133213221-0322103321120321-1123001031003030-2003001310002110"></a>

## Direct properties — DNS / 011002303000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022323311233030-3320300030032020-1112113121232031-0332333231002121-2322203313201012-0021320200323211-0033111130230222-0010133023133330"></a>

## Next pages — DNS / 011002303000 / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0023321322011323-3100133130131010-2130132310133012-0320103021203132-3302111230110322-3233011130100101-0303133131312311-3033112213223210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131302210013231-0333123123111322-1232033001102102-1120312111033103-2021330130133320-2300213101021200-0022122031330111-0013013221001013"></a>

## blocked_services.blocked_service.SSH — SSH / 300310100301 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032)
- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200)
- blocked_services.blocked_service.SSH

<a id="canonical-3103133302320113-2131131231122123-0010102212113001-2032203222102010-2302030112321200-2012213113131011-2020330211132100-1213330332013002"></a>

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
ssh = {}
```

<a id="canonical-0022100002011223-0312300002210110-2103232102010203-2223010122231312-2131233313120223-3222020111020320-0031331113023013-0131012322301021"></a>

## Direct properties — SSH / 300310100301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223100121001301-1200023122230010-0212313101223222-0011133112011023-2122330122331231-3231033021301020-2112221021322121-3312232231100320"></a>

## Next pages — SSH / 300310100301 / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3010321200301213-1221323211101322-2333032333112021-3301003333002322-1213032203210021-3210301202210213-3122033232331203-2020312120222230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211121012013311-2311112211321132-2222230222002033-2021001200102133-3012102011011301-3312132300330111-1310220310103332-3102110033021001"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 021332300322 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-2100213002102121-1020021110332332-3003003312033012-1320000102302111-0320002202323122-3121231313002122-2201111020312212-0032223101001032)
- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-0212221312333121-0120321230023130-0121033203122102-1331231121320122-2022001203203130-2011333233202333-3332002102203122-0230323122203320"></a>

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
web_user_interface = {}
```

<a id="canonical-2232233332032130-0223000002302113-1220200123321332-0121220032130220-2011030310321300-3023111231031032-0200111312131000-1310321030322321"></a>

## Direct properties — web_user_interface / 021332300322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211231201332030-3112111220332002-3210103223302112-1112020123131000-3323022010300220-1113123012131222-2020111220133321-1210331321332323"></a>

## Next pages — web_user_interface / 021332300322 / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-3022211222321313-1013032202102103-0232033020011203-0212303300100101-1103100210113233-0211330112110210-0211212200023130-1220033300212200)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1312001102001300-1311230131021121-2213102022222331-2331123300121313-1323121213000112-0102121102020303-2231311102230221-1322021220133331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011033120323333-1112111110010032-0123130112110312-3021121201001031-0301333102211032-1221111231112020-0022022210102323-2232120203323120"></a>

## cloud_credentials — cloud_credentials / 013131123121 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- cloud_credentials

<a id="canonical-2212230332101332-3103012112233333-1203033220233232-0101223001101320-2210013131021123-2003223030210221-0133303033131021-0201320211213022"></a>

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
cloud_credentials {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120320301133023-3231112301202213-0003321311100003-0333223130221301-2033102030330030-0123120320332011-1223220013220013-2110020320102022"></a>

## Direct properties — cloud_credentials / 013131123121 / 3

<a id="canonical-0210302331122110-0323000100132212-1300033130221332-1221003213212321-0033233001232333-1001200121233333-1323011133310231-0111210333033111"></a>

<a id="canonical-0010330130203232-0322330131231231-1012231111221212-2323001311132023-0111022001301232-3020131312303332-3100223222300012-0311020121121120"></a>

## name property — cloud_credentials / 013131123121 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1221300333301322-1320131302001102-3003103211100111-3211112231312110-3103101301331220-0222122330033331-1033233022312012-3203223213101011"></a>

<a id="canonical-1323030220211303-3121323323321002-1221333203323301-0311110200103303-1213020002100311-1200202102211103-2220032113033033-1223123131300031"></a>

## namespace property — cloud_credentials / 013131123121 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3101212133320031-0122203131232031-1211013120220033-1033231211332121-0321110203300322-1203121302023313-1230133301301031-3103022200331123"></a>

<a id="canonical-2303012332133222-1231101221102221-3133033233323300-1002011302031003-1232322331102313-3333002012212100-2320121111130130-0112213112222110"></a>

## tenant property — cloud_credentials / 013131123121 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0030030103310303-2221221010230202-3312313101021310-3103210111033203-3302103313103310-1113332320120310-0111333233023333-1021233001110101"></a>

## Next pages — cloud_credentials / 013131123121 / 7

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0132322102032221-0221031120301201-1222222320102200-3032133011203110-3012200020001320-2313113033220032-2222201130301012-2121000101223020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233031020113321-3223031231032111-0313031313110022-2211202020322232-0231310302321331-3332122302033020-3230113221130110-2320321312202230"></a>

## coordinates — coordinates / 130112020011 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- coordinates

<a id="canonical-1011322230220003-0302003212122331-3012223120323233-3131211303002303-0310020101322313-1023022223220102-3230322000001132-2203021002012121"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

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
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312323221101320-0303233032333010-0311303122012131-1133230301110222-3221132123012113-0200111013002100-1010221002131320-1013213120230312"></a>

## Direct properties — coordinates / 130112020011 / 3

<a id="canonical-0010200321313230-3131211010012101-1322003022301201-0111200320200221-3312313331312122-1310122102301002-1000221301102100-3221302000012100"></a>

<a id="canonical-3212132313112300-0033131102021322-3122200031122202-2203033100312323-2133111220220321-3201331103331122-1131210232100200-0100123213221111"></a>

## latitude property — coordinates / 130112020011 / 4

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-3123320120201112-3022203313103232-1220112011030001-1031323222221032-2301330212223022-3210100002231312-1233303132003111-0032223232322302"></a>

<a id="canonical-2201210133130220-1100302020312033-1212202222111311-2230010100112200-1130231211230231-0132131031322212-0301023021030130-1102012302300210"></a>

## longitude property — coordinates / 130112020011 / 5

Type: `"number"`. Optional.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-3211200202320002-3113130133303122-3233000221220200-3233331233023103-2301001031000131-2221130333211100-2230210233102323-0133302221023200"></a>

## Next pages — coordinates / 130112020011 / 6

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3220300023200020-3013222211030303-0100203211331300-0030113331233123-3331132002222231-0000301012011231-2310300112031020-2032323332321030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331321120123130-2210110100231332-0213000213132310-2112102302131211-3330101332101003-1220130002023201-2011302322022122-3232233210001213"></a>

## custom_dns — custom_dns / 222331112113 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- custom_dns

<a id="canonical-3002030232232100-2033020121311131-2133200130002313-1211330311002133-0010020103231120-0101101131302212-3001202113331320-2122033003121233"></a>

Type: `"object"`. single nested block, Optional.

Custom DNS is the configured for specify CE site.

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
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033131131323013-3132033033310313-3233210132022033-2200321301212010-2332232111300221-2212030032033123-3303111300102003-0022333233003130"></a>

## Direct properties — custom_dns / 222331112113 / 3

<a id="canonical-2232000032233212-2323112013020013-2201111330201213-0223011322133203-2300303233013122-3131033021212112-1123120123032130-3321321331013133"></a>

<a id="canonical-3210313101231220-3010313231223012-2111323101300211-0322212322211133-2101130032202100-2021223302222321-3032333022312310-0001133210222023"></a>

## inside_nameserver property — custom_dns / 222331112113 / 4

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2232030010333330-0012212102213311-0221200320033332-1220100022201320-3033331102012320-0121233011020100-1133321333233110-0201321021001030"></a>

<a id="canonical-0220000001323312-1200132113202332-3210122222311023-3032323032212012-1210110112001232-0300011122013010-1212320303233022-1103221122001332"></a>

## outside_nameserver property — custom_dns / 222331112113 / 5

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2032023031303223-3332021122020103-2002010223213012-1332210202301020-1121321010311120-1000333310020321-1111031323230220-0032120000233310"></a>

## Next pages — custom_dns / 222331112113 / 6

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3310031330331201-2123120101221211-0113032011130223-0003312130311301-1303121010101221-3233302131102330-2311002110111331-1013112312302133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322223323023200-3020110303023002-0131123320223210-1022101002023110-1111112010302231-1232222122113322-2301122332122223-1330312232331321"></a>

## default_blocked_services — default_blocked_services / 130003020013 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- default_blocked_services

<a id="canonical-3221320222002002-3132322013032031-1231302031321113-3320022220113332-3202301310311303-1213132233212112-2000011221113330-1331212333223123"></a>

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
default_blocked_services = {}
```

<a id="canonical-1221220121211131-1132232223100223-1212023202032230-1320023312023020-1032211032021011-2321113123121301-3000010300202120-0112203213113210"></a>

## Direct properties — default_blocked_services / 130003020013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111312112102122-3112122102113222-0000220100313213-0032200102200321-1230310011113311-3121111030130220-1030032311132130-3032322111302213"></a>

## Next pages — default_blocked_services / 130003020013 / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1131111322332132-3103020101000022-2211033321232233-2030233213233230-0131222130003033-0220320200112322-2313123101130311-0133321200322303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232101012111012-3333121132333130-1330312101232223-2032003131232003-2311112102022120-3312310212123010-3000020122110102-2310330331222032"></a>

## disable_encryption — disable_encryption / 220112310133 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- disable_encryption

<a id="canonical-1211200031130233-1323333201331301-0322231221101222-1120210211030120-0120010001323132-1212132221021213-0323103002020013-1303120013230211"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

OneOf alternatives in this subsection:

- [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1211200031130233-1323333201331301-0322231221101222-1120210211030120-0120010001323132-1212132221021213-0323103002020013-1303120013230211)
- [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-1201103122112231-2220103102133220-0020113032201213-2220030020301213-0121101300223203-1022230101003312-3002300313221133-2301021211022321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_encryption = {}
```

<a id="canonical-1300112323321112-2231002322133103-3100203000023322-0033201301001310-1203011101333203-2323032021301000-0322213032121333-0330101020032101"></a>

## Direct properties — disable_encryption / 220112310133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002103130313013-3033230203213012-1320203230121221-0220131012022222-3000111303000131-3130033110320302-1211321323310320-1030331212313121"></a>

## Next pages — disable_encryption / 220112310133 / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1101221123132310-1010130003200313-0032320202321131-0203333233313013-2011212112020022-2213003211113113-1032113120210301-2132310133203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310312001001331-2231120130211100-3200200332001322-2130212020303202-2011112211120001-1032210202223032-1102012320202222-2010231320133313"></a>

## enable_encryption — enable_encryption / 213301112033 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- enable_encryption

<a id="canonical-1201103122112231-2220103102133220-0020113032201213-2220030020301213-0121101300223203-1022230101003312-3002300313221133-2301021211022321"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("kms_key_resource_id",
    "kms_key_ring_id")}
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
enable_encryption {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230023021032332-1303121023232022-1032123000310213-0312220213312330-0301130103201031-0113101310332023-0332323212111113-1223321303101011"></a>

## Direct properties — enable_encryption / 213301112033 / 3

<a id="canonical-1331133121310212-1132003312013130-0230222310313103-2121122323221302-3000132331102322-3122022131002011-2033001012030110-3112020302202110"></a>

<a id="canonical-3232331002300230-2103131012132300-2130212211120303-1232110202331333-3110301122003022-0200303332132211-3032332210000132-2101022100222323"></a>

## kms_key_resource_id property — enable_encryption / 213301112033 / 4

Type: `"string"`. Optional.

GCP KMS Key to be used to encrypt the disk attached to the VM.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2333220201230320-1023313303133030-0031220230301211-3111111003033122-0110213200323301-3303202002311232-1313200023322233-2023211202130000"></a>

<a id="canonical-3311022121113023-0211000301101331-2302003221230202-1230303031003031-0122312231101222-2203321100220330-3202113312230232-1320010323223030"></a>

## kms_key_ring_id property — enable_encryption / 213301112033 / 5

Type: `"string"`. Optional.

Key ring in which the CMK to be used to encrypt is present.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230101100101211-3223023022033212-3212100302202031-2222011210221103-0233113301310310-0313003033002333-2030012012020223-2022303123311321"></a>

## Next pages — enable_encryption / 213301112033 / 6

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323233110033110-1203010023000231-0202322320222001-0320203110310111-0013210310121310-3332222211120002-1132311000221322-3002030130120000"></a>

## ingress_egress_gw — ingress_egress_gw / 331213233023 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- ingress_egress_gw

<a id="canonical-3012232203320333-1330203033310313-3102123222033211-1333321322230123-1322023301131002-1212320311301032-0121011000200030-3013011002010022"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface GCP ingress/egress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("gcp_certified_hw",
    "gcp_zone_names"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3012232203320333-1330203033310313-3102123222033211-1333321322230123-1322023301131002-1212320311301032-0121011000200030-3013011002010022)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-2121210113102303-2011303133203313-3210030103333203-3020313020302201-0111103112101112-3110132101132312-3133132320300133-3310212103131122)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-1312003312232212-2011113310110131-0010131322113332-3131330000201333-3320312120210223-3032202123102121-1200233120130210-3103303212322322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ingress_egress_gw {
  # Configure direct properties listed below.
}
```
