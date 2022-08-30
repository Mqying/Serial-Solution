import { TableType } from "@/tableType.js";
import { actualOutput } from "./actualOutput";
import { formatTime, formatUnit } from "../../services/exportRecords/format";
import zhCN from "@/locales/zh-CN/pages";
import enUS from "@/locales/en-US/pages";

export function exportRecords(selectedRowsState, selectType, locale) {
    const language = locale == 'zh-CN' ? zhCN : enUS

    let head;
    let columns = [];

    switch (selectType) {
        case TableType.dielectric:
            head = language["pages.table.export.head.dielectric"];

            columns = new Array(13);

            for (let i = 0; i < columns.length; i++) {
                columns[i] = new Array();
            }

            columns[0].push(language["pages.table.export.head.index"]);

            columns[1].push(language["pages.table.detectionTime"]);
            columns[2].push(language["pages.table.avg"]);

            for (let i = 3; i < columns.length; i++) {
                columns[i].push(`${language["pages.table.frequency"]} ${i - 2}`);
            }

            for (let i = 0; i < selectedRowsState.length; i++) {
                columns[0].push(selectedRowsState[i].index);
                columns[1].push(formatTime(selectedRowsState[i].detectionTime));
                columns[2].push(`${selectedRowsState[i].average}`);

                for (let k = 0; k < selectedRowsState[i].items.length; k++) {
                    columns[k + 3].push(`${selectedRowsState[i].items[k]} kV`);
                }
            }

            break;

        case TableType.water:
            head = language["pages.table.export.head.water"];

            columns = new Array(4);

            for (let i = 0; i < columns.length; i++) {
                columns[i] = new Array();
            }

            columns[0].push(language["pages.table.detectionTime"]);
            columns[1].push(language["pages.table.quantity"]);
            columns[2].push(language["pages.table.percentage"]);
            columns[3].push(language["pages.table.percentage"]);

            for (let i = 0; i < selectedRowsState.length; i++) {
                columns[0].push(formatTime(selectedRowsState[i].detectionTime));
                columns[1].push(`${selectedRowsState[i].quantity}`);
                columns[2].push(`${selectedRowsState[i].ratio1}`);
                columns[3].push(`${selectedRowsState[i].ratio2}`);
            }

            break;

        case TableType.acid:
            head = language["pages.table.export.head.acid"];

            columns = new Array(7);

            for (let i = 0; i < columns.length; i++) {
                columns[i] = new Array();
            }

            columns[0].push(language["pages.table.detectionTime"]);

            for (let i = 1; i < columns.length; i++) {
                columns[i].push(`PH ${i}`);
            }

            for (let i = 0; i < selectedRowsState.length; i++) {
                columns[0].push(formatTime(selectedRowsState[i].detectionTime));

                for (let k = 0; k < selectedRowsState[i].items.length; k++) {
                    columns[k + 1].push(`${selectedRowsState[i].items[k]}`);
                }
            }

            break;

        case TableType.flash:
            head = language["pages.table.export.head.flash"];

            columns = new Array(5);

            for (let i = 0; i < columns.length; i++) {
                columns[i] = new Array();
            }

            columns[0].push(language["pages.table.detectionTime"]);
            columns[1].push(language["pages.table.atmosphericPressure"]);
            columns[2].push(language["pages.table.preFlashTemperature"]);
            columns[3].push(language["pages.table.flashPointTemperature"]);
            columns[4].push(language["pages.table.sampleSerialNumber"]);

            for (let i = 0; i < selectedRowsState.length; i++) {
                columns[0].push(formatTime(selectedRowsState[i].detectionTime));
                columns[1].push(`${selectedRowsState[i].pressure}`);
                columns[2].push(`${selectedRowsState[i].preTemperature}`);
                columns[3].push(`${selectedRowsState[i].pointTemperature}`);
                columns[4].push(`${selectedRowsState[i].sampleNumber}`);
            }

            break;
    }

    actualOutput({
        head: head,
        columns: columns
    });
}