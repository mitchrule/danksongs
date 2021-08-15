import { Button } from "@material-ui/core";
import { Formik, Field, Form } from "formik";
import { Container } from "next/app";

const API_URL = '/';

// @TODO assign submitData type

const handleSubmit = async (submitData) => {
    const res = await fetch(API_URL, submitData);
}

export default function LoginForm() {
    return (
        <Container>
            <Formik
                 initialValues={{ name: "", email: "" }}
                 onSubmit={async (submitData) => handleSubmit(submitData)}
            >
                <Form>
                    <Field name="username" type="text" />
                    <Field name="password" type="password"/>
                    <Button>Sign In</Button>
                </Form>
            </Formik>

        </Container>
        "TODO LoginForm"
    
}