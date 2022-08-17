import React from "react";

// See React-Bootstrap Alert component for help.
export interface GlobalAlertProperties {
    show: boolean;
    content: string;
    variant: "success" | "primary" | "danger" | "warning" | "info" | "secondary" | "light" | "dark";
}
