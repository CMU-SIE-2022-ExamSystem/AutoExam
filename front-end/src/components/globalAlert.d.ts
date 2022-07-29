import React from "react";

export interface GlobalAlertProperties {
    show: boolean;
    content: string;
    variant: "success" | "primary" | "danger" | "warning" | "info" | "secondary" | "light" | "dark";
}
