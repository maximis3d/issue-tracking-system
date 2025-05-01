import React from "react";
import * as XLSX from "xlsx";
import { saveAs } from "file-saver";

const ExcelExportButton = ({ data, fileName = "export", sheetName = "Sheet1" }) => {
  const handleExport = () => {
    if (!data || data.length === 0) return;

    const worksheet = XLSX.utils.json_to_sheet(data);
    const workbook = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(workbook, worksheet, sheetName);

    const excelBuffer = XLSX.write(workbook, { bookType: "xlsx", type: "array" });
    const blob = new Blob([excelBuffer], { type: "application/octet-stream" });
    saveAs(blob, `${fileName}.xlsx`);
  };

  return (
    <button
      onClick={handleExport}
      className="bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 px-6 rounded-md shadow-lg focus:outline-none focus:ring-2 focus:ring-blue-400 transition duration-300 ease-in-out"
    >
      Export to Excel
    </button>
  );
};

export default ExcelExportButton;
