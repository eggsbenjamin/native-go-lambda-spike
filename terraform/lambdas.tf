resource "aws_lambda_function" "invoker" {
  filename         = "../artifacts/${var.tag}invoker.zip"
  function_name    = "${var.tag}invoker"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "${var.tag}invoker"
  source_code_hash = "${base64sha256(file("../artifacts/${var.tag}invoker.zip"))}"
  runtime          = "go1.x"

  environment {
    variables = {
      STEP_FN_ARN = "${aws_sfn_state_machine.demo_step_function.id}"
    }
  }
}

resource "aws_lambda_function" "age_check" {
  filename         = "../artifacts/${var.tag}age_check.zip"
  function_name    = "${var.tag}age_check"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "${var.tag}age_check"
  source_code_hash = "${base64sha256(file("../artifacts/${var.tag}age_check.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "check_hair_colour" {
  filename         = "../artifacts/${var.tag}check_hair_colour.zip"
  function_name    = "${var.tag}check_hair_colour"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "${var.tag}check_hair_colour"
  source_code_hash = "${base64sha256(file("../artifacts/${var.tag}check_hair_colour.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "has_ginger_friends" {
  filename         = "../artifacts/${var.tag}has_ginger_friends.zip"
  function_name    = "${var.tag}has_ginger_friends"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "${var.tag}has_ginger_friends"
  source_code_hash = "${base64sha256(file("../artifacts/${var.tag}has_ginger_friends.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "cool" {
  filename         = "../artifacts/${var.tag}cool.zip"
  function_name    = "${var.tag}cool"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "${var.tag}cool"
  source_code_hash = "${base64sha256(file("../artifacts/${var.tag}cool.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "not_cool" {
  filename         = "../artifacts/${var.tag}not_cool.zip"
  function_name    = "${var.tag}not_cool"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "${var.tag}not_cool"
  source_code_hash = "${base64sha256(file("../artifacts/${var.tag}not_cool.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "under_age" {
  filename         = "../artifacts/${var.tag}under_age.zip"
  function_name    = "${var.tag}under_age"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "${var.tag}under_age"
  source_code_hash = "${base64sha256(file("../artifacts/${var.tag}under_age.zip"))}"
  runtime          = "go1.x"
}
