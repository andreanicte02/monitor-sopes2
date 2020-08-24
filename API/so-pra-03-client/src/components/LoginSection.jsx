import React from 'react';


import {Card, Icon, message} from "antd";

import * as Yup from "yup";
import {Formik} from "formik";
import {Form, Input, InputNumber, ResetButton, SubmitButton} from "@jbuschke/formik-antd";

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

// onLoginSuccess
export default class LoginSection extends React.Component {


    onFormSubmit = async (values, {setSubmitting, resetForm}) => {
        console.log('los valores son ');
        console.log(values);

        await sleep(500);

        if (values.username === 'admin' && values.password === 'admin') {
            this.props.onLoginSuccess();
        } else {
            message.error('Credenciales incorrectas');
        }

        setSubmitting(false);
    };



    render() {

        const validationSchema = Yup.object().shape({
            username: Yup.string()
                .required('Username es requerido'),
            password: Yup.string()
                .required('Password es requerido')
        });

        return (<div style={{ }}>
            <Card title="Login" style={{width: 800, margin: "50px auto" }}>

            <Formik
                validationSchema={validationSchema}
                onSubmit={this.onFormSubmit}
                initialValues={ {username: "", password: ""} }
            >
                {() => (
                    <Form layout="inline" >

                        <Form.Item label="Username" name="username">
                            <Input
                                prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                                name="username"
                                placeholder="Username" />
                        </Form.Item>

                        <Form.Item label="Password" name="password">
                            <Input
                                prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                                name="password"
                                type="password"
                                placeholder="Password" />
                        </Form.Item>

                        <SubmitButton>Log in</SubmitButton>
                    </Form>
                )}

            </Formik>

        </Card>
        </div>);
    }
}