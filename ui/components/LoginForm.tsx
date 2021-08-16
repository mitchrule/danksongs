import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Formik, Field, Form, FormikHelpers } from "formik";
import { Container } from "next/app";
import { Box, Button, Typography } from '@material-ui/core';
import { useState } from 'react';

const API_URL = '/';
const FORM_DEFAULT = 'sign-in';


interface SignInValues {
    userName: string;
    password: string;
  }

interface SignUpValues {
    userName: string;
    password1: string;
    password2: string;
  }
  

export default function LoginForm() {
    const [formScreen, setFormScreen] = useState(FORM_DEFAULT);

    const handleSignIn = async ({}) => {
        // TODO
        const res = await fetch(API_URL);
    }
    
    const handleSignUp = async ({}) => {
        // TODO
        const res = await fetch(API_URL);
    } 
    
    const handleClick = () => {
        if (formScreen === 'sign-up') {
            setFormScreen('sign-in');
        } else {
            setFormScreen('sign-up');
        }
    }

    return (   
        <Container> 
        {/** By default let the user sign in */}
        {formScreen === 'sign-in' && (
            <Formik
                initialValues={{
                userName: '',
                password: '',
                }}
                onSubmit={(
                values: SignInValues,
                { setSubmitting }: FormikHelpers <SignInValues>
                ) => {
                setTimeout(() => {
                    handleSignIn(values);
                    setSubmitting(false);
                }, 500);
                }}
            >
                <Form>
                <Typography>Username</Typography>
                <Field id="userName" name="userName" placeholder="test" />

                <Typography>Password</Typography>
                <Field
                    id="email"
                    name="email"
                    placeholder="password"
                    type="password"
                />
                <Button type="submit">Log In</Button>
                </Form>
            </Formik>
        )}

        {/** Otherwise preform the sign up here */}
        {formScreen === 'sign-up' && (
            <Formik
            initialValues={{
            userName: '',
            password1: '',
            password2: '',
            }}
            onSubmit={(
            values: SignUpValues,
            { setSubmitting }: FormikHelpers <SignUpValues>
            ) => {
            setTimeout(() => {
                handleSignUp(values);
                setSubmitting(false);
            }, 500);
            }}
        >
            <Form>
            <Typography>Username</Typography>
            <Field id="userName" name="userName" placeholder="test" />

            <Typography>Password</Typography>
            <Field
                id="password1"
                name="password1"
                placeholder="password"
                type="password"
            />
            <Typography>Re-Enter Password</Typography>
            <Field
                id="password2"
                name="password2"
                placeholder="password again"
                type="password"
            />
            <Button type="submit">Sign Up!</Button>
            </Form>
            </Formik>
        )}
        {/* Control form display */}
        <Box>
            <Button onClick={handleClick}>
                {formScreen === 'sign-up' ? 'Sign In Instead' : 'Sign Up Instead' }
            </Button>
        </Box>
        </Container>
    );
}