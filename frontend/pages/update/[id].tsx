import type { NextPage, NextPageContext } from "next";
import { Button, useToast } from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import HookFormField from "../../components/HookFormField";
import { useRouter } from 'next/router'

const putReq = async (
    url: string,
    data: Record<string, string | number>
  ): Promise<{ success: boolean }> => {
    data["Quantity"] = +data["Quantity"]
    const response = await (
      await fetch(url, {
        method: "PUT",
        mode: "cors",
        body: JSON.stringify(data)
      })
    ).json();
  
    return response;
  };

const Update: NextPage = () => {
    const { handleSubmit, register, formState } = useForm();
    const { isSubmitting } = formState;
    const toast = useToast();
    const router = useRouter()
    const { id } = router.query

    const onSubmit = async (values: Record<string, string | number>) => {
      putReq(
          `https://inventory-tracking-uk.herokuapp.com/update/${id}`,
          values
        ).then(function(success) {
          toast({
            title: success ? "Item Successfully Updated" : "Item Updation Failed",
            status: success ? "success" : "error",
            duration: 3000,
            isClosable: true,
            onCloseComplete: () => window.location.href = "/",
          });
        })
    };
    

    return (
      <form onSubmit={handleSubmit(onSubmit)}>
        <HookFormField
          name="Name"
          label="Name"
          formState={formState}
          register={register}
          registerOptions={{
            minLength: { value: 1, message: "Minimum length should be 1" },
            maxLength: { value: 30, message: "Maximum length should be 30" },
          }}
          type="text"
        />
  
        <HookFormField
          name="Quantity"
          label="Quantity"
          formState={formState}
          register={register}
          registerOptions={{
            min: { value: 0, message: "Minimum quantity should be 0" },
          }}
          type="number"
        />

        <HookFormField
          name="Location"
          label="Location"
          formState={formState}
          register={register}
          registerOptions={{}}
          type="text"
        />
        <Button mt={4} colorScheme="green" isLoading={isSubmitting} type="submit">
          Update
        </Button>
      </form>
    );
  };

export default Update;