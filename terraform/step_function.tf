resource "aws_sfn_state_machine" "demo_step_function" {
  name     = "demo_demo_step_function"
  role_arn = "${aws_iam_role.demo_step_function.arn}"

  definition = "${data.template_file.demo_step_function_config.rendered}"
}

data "template_file" "demo_step_function_config" {
  template = "${file("${path.module}/demo_step_function.json")}"

  vars {
    age_check_arn         = "${aws_lambda_function.age_check.arn}"
    check_hair_colour_arn = "${aws_lambda_function.check_hair_colour.arn}"
    has_ginger_friends    = "${aws_lambda_function.has_ginger_friends.arn}"
    under_age_arn         = "${aws_lambda_function.under_age.arn}"
    cool_arn              = "${aws_lambda_function.cool.arn}"
    not_cool_arn          = "${aws_lambda_function.not_cool.arn}"
  }
}
