---
layout: "aws"
page_title: "AWS: aws_iam_service_linked_role"
sidebar_current: "docs-aws-resource-iam-service-linked-role"
description: |-
  Provides an IAM service-linked role.
---

# aws_iam_service_linked_role

Provides an [IAM service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html).

## Example Usage

```hcl
resource "aws_iam_service_linked_role" "elasticbeanstalk" {
  aws_service_name = "elasticbeanstalk.amazonaws.com"
}
```

## Argument Reference

The following arguments are supported:

* `aws_service_name` - (Required, Forces new resource) The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).

## Attributes Reference

The following attributes are exported:

* `name` - The name of the role.
* `path` - The path of the role.
* `arn` - The Amazon Resource Name (ARN) specifying the role.
* `create_date` - The creation date of the IAM role.
* `unique_id` - The stable and unique string identifying the role.
* `description` - The description of the role.

## Import

IAM service-linked roles can be imported using role ARN, e.g.

```
$ terraform import aws_iam_service_linked_role.elasticbeanstalk arn:aws:iam::123456789012:role/aws-service-role/elasticbeanstalk.amazonaws.com/AWSServiceRoleForElasticBeanstalk
```
