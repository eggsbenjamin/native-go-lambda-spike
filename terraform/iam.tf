resource "aws_iam_role_policy" "lambda_policy" {
  name = "lambda_policy"
  role = "${aws_iam_role.lambda_role.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*",
        "ec2:CreateNetworkInterface",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DeleteNetworkInterface",
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "states:StartExecution",
        "states:DescribeExecution"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role" "demo_step_function" {
  name = "demo_step_function_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "states.eu-west-1.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "demo_step_function" {
  name   = "demo_step_function_policy"
  role   = "${aws_iam_role.demo_step_function.name}"
  policy = "${data.aws_iam_policy_document.allow_invoke_lambda.json}"
}

data "aws_iam_policy_document" "allow_invoke_lambda" {
  statement {
    actions = [
      "lambda:InvokeFunction",
    ]

    resources = [
      "${aws_lambda_function.age_check.arn}",
      "${aws_lambda_function.check_hair_colour.arn}",
      "${aws_lambda_function.has_ginger_friends.arn}",
      "${aws_lambda_function.under_age.arn}",
      "${aws_lambda_function.cool.arn}",
      "${aws_lambda_function.not_cool.arn}"
    ]
  }
}
