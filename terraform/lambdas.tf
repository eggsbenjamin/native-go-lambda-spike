resource "aws_lambda_function" "invoker" {
  filename         = "../artifacts/invoker.zip"
  function_name    = "invoker"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "invoker"
  source_code_hash = "${base64sha256(file("../artifacts/invoker.zip"))}"
  runtime          = "go1.x"

  environment {
    variables = {
      STEP_FN_ARN = "${aws_sfn_state_machine.demo_step_function.id}"
    }
  }
}

resource "aws_lambda_function" "age_check" {
  filename         = "../artifacts/age_check.zip"
  function_name    = "age_check"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "age_check"
  source_code_hash = "${base64sha256(file("../artifacts/age_check.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "check_hair_colour" {
  filename         = "../artifacts/check_hair_colour.zip"
  function_name    = "check_hair_colour"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "check_hair_colour"
  source_code_hash = "${base64sha256(file("../artifacts/check_hair_colour.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "has_ginger_friends" {
  filename         = "../artifacts/has_ginger_friends.zip"
  function_name    = "has_ginger_friends"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "has_ginger_friends"
  source_code_hash = "${base64sha256(file("../artifacts/has_ginger_friends.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "cool" {
  filename         = "../artifacts/cool.zip"
  function_name    = "cool"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "cool"
  source_code_hash = "${base64sha256(file("../artifacts/cool.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "not_cool" {
  filename         = "../artifacts/not_cool.zip"
  function_name    = "not_cool"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "not_cool"
  source_code_hash = "${base64sha256(file("../artifacts/not_cool.zip"))}"
  runtime          = "go1.x"
}

resource "aws_lambda_function" "under_age" {
  filename         = "../artifacts/under_age.zip"
  function_name    = "under_age"
  role             = "${aws_iam_role.lambda_role.arn}"
  handler          = "under_age"
  source_code_hash = "${base64sha256(file("../artifacts/under_age.zip"))}"
  runtime          = "go1.x"
}
