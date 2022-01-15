import type {
    FormState,
    FieldValues,
    UseFormRegister,
    Message,
    ValidationRule,
  } from "react-hook-form";
  import {
    FormErrorMessage,
    FormLabel,
    FormControl,
    Input,
  } from "@chakra-ui/react";
  
  export type RegisterOptions = Partial<{
    required: Message | ValidationRule<boolean>;
    min: ValidationRule<number>;
    max: ValidationRule<number>;
    maxLength: ValidationRule<number>;
    minLength: ValidationRule<number>;
  }>;
  
  const HookFormField = ({
    name,
    label,
    formState,
    register,
    registerOptions,
    type,
  }: {
    name: string;
    label: string;
    formState: FormState<FieldValues>;
    register: UseFormRegister<FieldValues>;
    registerOptions: RegisterOptions;
    type: "text" | "number";
  }) => {
    const { errors } = formState;
  
    return (
      <FormControl isInvalid={errors[name]}>
        <FormLabel htmlFor={name}>{label}</FormLabel>
        <Input
          id={name}
          placeholder={label}
          type={type}
          {...register(name, registerOptions)}
        />
        <FormErrorMessage>
          {errors[name] && errors[name].message}
        </FormErrorMessage>
      </FormControl>
    );
  };
  
  export default HookFormField;