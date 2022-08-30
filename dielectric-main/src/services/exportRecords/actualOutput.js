import ExcelJS from "exceljs";

export async function actualOutput(data) {
    let head = data.head || false;
    let columns = data.columns;

    let workbook = new ExcelJS.Workbook();
    let worksheet = workbook.addWorksheet("Data Table");

    for (let i = 1; i <= columns.length; i++) {
        worksheet.getColumn(i).values = columns[i - 1];
    }

    let border = {
        top: { style: 'thin' },
        left: { style: 'thin' },
        bottom: { style: 'thin' },
        right: { style: 'thin' }
    };

    let defaultAlignment = {
        vertical: 'middle',
        horizontal: 'center'
    };

    let defaultRowHeight = 26;

    worksheet.eachRow(
        (row) => {
            row.height = defaultRowHeight

            row.eachCell(
                (cell) => {
                    cell.border = border;
                    cell.alignment = defaultAlignment;
                }
            )
        }
    );

    for (let col = 0; col < columns.length; col++) {
        let length = -1;

        for (let row = 0; row < columns[col].length; row++) {
            let currentCellLength = (() => {
                let cellContent = columns[col][row].toString();
                let result = 0;

                for (let i = 0; i < cellContent.length; i++) {
                    if (cellContent.charCodeAt(i) > 127) {
                        result += 2;
                    } else {
                        result++;
                    }
                }

                return result;
            })();

            if (currentCellLength > length) {
                length = currentCellLength;
            }
        }

        length = length < 15 ? 15 : length + 2;

        worksheet.getColumn(col + 1).width = length;
    }


    if (head) {
        worksheet.insertRow(1, head);

        worksheet.mergeCells(
            `A1:${String.fromCharCode('A'.charCodeAt() + data.columns.length - 1)}1`
        );

        worksheet.getCell("A1").value = head;

        let headFont = {
            size: 12,
            bold: true
        };

        worksheet.getCell("A1").font = headFont;
        worksheet.getCell("A1").alignment = defaultAlignment;
        worksheet.getCell("A1").height = defaultRowHeight
    }

    workbook.xlsx.writeBuffer().then((buffer) => {
        let blob = new Blob([buffer], {
            type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
        });
        let url = window.URL.createObjectURL(blob);
        let anchor = document.createElement('a', {});

        anchor.href = url;

        let documentName = 'historyData';
        anchor.download = `${documentName}.xlsx`;

        anchor.click();

        window.URL.revokeObjectURL(url);
    })
}